package middleware

import (
	"github.com/kenSevLeb/go-framework/component/rbac"
	"github.com/kenSevLeb/go-framework/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Permission 权限校验中间件
func Permission(permission *rbac.Permission) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取UID和domain
		uid := ctx.GetString("user_id")

		// 根据用户ID,路由，请求方法校验权限
		hasPermission := permission.MustValidate(uid, ctx.Request.URL.Path)

		// 没有权限
		if !hasPermission {
			response.Wrap(ctx).Error(http.StatusUnauthorized, "StatusUnauthorized")
			ctx.Abort()
		}

		ctx.Next()
	}

}
