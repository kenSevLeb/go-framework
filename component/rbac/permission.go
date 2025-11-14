package rbac

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

// RolePermission 角色权限
type RolePermission struct {
	// 角色ID
	RoleId string
	// 路由
	Route string
	// 域名,暂未使用
	Domain string
}

// UserRole 用户角色
type UserRole struct {
	// 用户ID
	UserId string
	// 角色ID
	RoleId string
}

// 权限
type Permission struct {
	// enforcer实例
	enforcer *casbin.Enforcer
	// 选项
	opts options
}

// 实例化
func New() *Permission {
	return new(Permission)
}

// Load 设置规则，使用时必须实现该函数
func (p *Permission) Load(
	rolePermissionProvider func() []RolePermission, // 获取角色权限
	userRoleProvider func() []UserRole, // 获取用户角色
	opts ...Option, // 选项
) error {
	p.opts.rolePrefix = "goRole" // 默认值
	// 设置options
	for _, opt := range opts {
		opt.apply(&p.opts)
	}
	// 初始化适配器
	adapter := &Adapter{
		getRolePermissionsFunc: rolePermissionProvider,
		getUserRolesFunc:       userRoleProvider,
		rolePrefix:             p.opts.rolePrefix,
	}

	var err error
	p.enforcer, err = adapter.GetEnforcer()
	if err != nil { // print err but not exit
		return fmt.Errorf("get enforcer: %v", err)
	}

	// 加载分布式watcher
	if p.opts.watcher != nil {
		if err := p.enforcer.SetWatcher(p.opts.watcher); err != nil {
			return fmt.Errorf("set watcher: %v", err)
		}
	}

	return nil
}

// 是否为超管
func (p *Permission) isAdministrator(sub string) bool {
	if p.opts.adminRole == "" {
		return false
	}
	if ok, _ := p.enforcer.HasRoleForUser(sub, roleWithPrefix(p.opts.rolePrefix, p.opts.adminRole)); ok {
		return true
	}
	return false
}

// 验证权限
func (p *Permission) Validate(sub string, obj string) (bool, error) {
	// 检查是否为超管
	if p.isAdministrator(sub) {
		return true, nil
	}

	result, err := p.enforcer.Enforce(sub, obj, p.opts.domain)
	if result == true {
		return true, nil
	}

	// 二次匹配的条件，比如某某系统，对于未能匹配的路由，也要允许访问
	if p.opts.pathFilter != nil && p.opts.pathFilter(p.enforcer.GetAllObjects(), obj) {
		return true, nil
	}

	return false, err
}

// 验证权限，只返回bool
//
// sub为用户标识,obj为路由
func (p *Permission) MustValidate(sub string, obj string) bool {
	has, _ := p.Validate(sub, obj)
	return has
}

// Sync 同步缓存,在更新DB成功后调用
func (p *Permission) Sync() error {
	return p.enforcer.SavePolicy()
}
