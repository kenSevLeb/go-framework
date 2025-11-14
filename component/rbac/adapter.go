package rbac

import (
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"log"
)

const (
	TypeResource = "p" // 权限类型
	TypeRole     = "g" // 角色类型
)

var (
	ErrInvalidOperate = errors.New("invalid operate")
)

// 自定义适配器
type Adapter struct {
	// 提供角色权限数据的方法
	getRolePermissionsFunc func() []RolePermission
	// 提供用户角色数据的方法
	getUserRolesFunc func() []UserRole
	// role prefix
	rolePrefix string
}

// LoadPolicy 加载权限规则
func (a *Adapter) LoadPolicy(model model.Model) error {
	log.Println("load policy")
	// 加载新规则
	rolePermissions := a.getRolePermissionsFunc()
	for _, item := range rolePermissions {
		// 获取相关权限并加载到rbac里
		// 角色ID，路由
		lineText := fmt.Sprintf("%s,%s,%s,%s", TypeResource, roleWithPrefix(a.rolePrefix, item.RoleId), item.Route, item.Domain)
		persist.LoadPolicyLine(lineText, model)
	}

	userRoles := a.getUserRolesFunc()
	for _, item := range userRoles {
		// 获取相关权限并加载到rbac里
		// 用户ID，角色ID
		lineText := fmt.Sprintf("%s,%s,%s", TypeRole, item.UserId, roleWithPrefix(a.rolePrefix, item.RoleId))
		persist.LoadPolicyLine(lineText, model)
	}

	return nil
}

// SavePolicy 保存权限规则
func (a Adapter) SavePolicy(model model.Model) error {
	return a.LoadPolicy(model)
}

// AddPoliy 添加权限规则
func (a Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	return ErrInvalidOperate
}

// 删除权限规则
func (a Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return ErrInvalidOperate
}

// RemoveFilteredPolicy
func (a Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return ErrInvalidOperate
}

// rbacModel 定义基于角色的rbac模型
const rbacModel = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
# 支持正则
m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act)
`

// GetEnforcer 基于自定义适配器的权限管理器
func (a *Adapter) GetEnforcer() (*casbin.Enforcer, error) {
	// 默认
	enforcer := new(casbin.Enforcer)
	m, err := model.NewModelFromString(rbacModel)
	if err != nil {
		return enforcer, err
	}
	authEnforcer, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return enforcer, err
	}
	return authEnforcer, nil
}
