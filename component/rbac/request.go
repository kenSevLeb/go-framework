package rbac

type PermissionSyncRequest struct {
	Items []PermissionItem
}

type PermissionItem struct {
	Path      string `json:"path" form:"path"`
	Type      string `json:"type" form:"type"`
	Name      string `json:"title" form:"title"`
	UniqueKey string `json:"key" form:"key"`
	Routers   []Router
}

// 路由
type Router struct {
	Route string
}

// 绑定角色权限
type BindRolePermissionRequest struct {
	// 角色名称
	RoleName string `json:"role_name" form:"role_name" binding:"required" comment:"角色名称"`
	// 角色ID，创建时为0,更新时不为0
	RoleId int `json:"role_id" form:"role_id" comment:"角色ID"`
	// 已逗号连接的key
	UniqueKeys string `json:"unique_keys" form:"unique_keys" comment:"唯一key"`
}

// 绑定用户角色
type BindUserRoleRequest struct {
	// 用户ID
	UserId int `json:"user_id" form:"user_id" comment:"用户ID"`
	// 角色ID，已逗号隔开
	RoleIds string `json:"role_ids" form:"role_ids" comment:"角色ID"`
}
