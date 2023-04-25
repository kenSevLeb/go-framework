package rbac

import (
	"github.com/gin-gonic/gin"
)

// Option 选项
type HandlerOption struct {
	Prefix                string
	IndexHandler          gin.HandlerFunc // 权限树
	SyncHandler           gin.HandlerFunc // 同步权限树
	ListRoleHandler       gin.HandlerFunc // 角色列表
	CreateRoleHandler     gin.HandlerFunc // 创建角色
	DeleteRoleHandler     gin.HandlerFunc // 删除角色
	BindPermissionHandler gin.HandlerFunc // 给角色绑定权限
	BindRoleHandler       gin.HandlerFunc // 给用户绑定角色

}

// Register 使用默认配置注册路由
func (p *Permission) Register(r *gin.Engine) {
	p.RouterRegister(&(r.RouterGroup), HandlerOption{Prefix: "rbac"})
}

// RouterRegister 注册路由
func (p *Permission) RouterRegister(rg *gin.RouterGroup, option HandlerOption) {
	prefixGroup := rg.Group(option.Prefix)
	{
		prefixGroup.GET("index", useDefault(option.IndexHandler, p.indexHandler()))
		prefixGroup.GET("sync", useDefault(option.SyncHandler, p.syncHandler()))
		prefixGroup.GET("bindRouter", useDefault(option.BindRoleHandler, p.bindRoleHandler()))
		prefixGroup.GET("listRole", useDefault(option.ListRoleHandler, p.listRoleHandler()))
		prefixGroup.GET("createRole", useDefault(option.CreateRoleHandler, p.createRoleHandler()))
		prefixGroup.GET("deleteRole", useDefault(option.DeleteRoleHandler, p.deleteRoleHandler()))
		prefixGroup.GET("bindPermission", useDefault(option.BindPermissionHandler, p.bindPermissionHandler()))
		prefixGroup.GET("bindRole", useDefault(option.BindRoleHandler, p.bindRoleHandler()))
	}
}

func useDefault(needHandler, defaultHandler gin.HandlerFunc) gin.HandlerFunc {
	if needHandler != nil {
		return needHandler
	}
	return defaultHandler
}

// 权限树列表
func (p *Permission) indexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// 同步权限
func (p *Permission) syncHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (p *Permission) listRoleHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (p *Permission) createRoleHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (p *Permission) deleteRoleHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (p *Permission) bindPermissionHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func (p *Permission) bindRoleHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
