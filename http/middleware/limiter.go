package middleware

import (
	"git.yingxiong.com/platform/go-framework/component/limiter"
	"git.yingxiong.com/platform/go-framework/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Limiter 限流器
func Limiter(limitValue int) gin.HandlerFunc {
	limit := limiter.MaxAllowPerSecond(limitValue)
	return func(ctx *gin.Context) {

		if !limit.Allow() {
			response.Wrap(ctx).Error(http.StatusServiceUnavailable, "StatusServiceUnavailable")
			ctx.Abort()
		}
		ctx.Next()
	}
}

// LimiterWithUri 根据URI限流器
func LimiterWithUri(limitValue int) gin.HandlerFunc {
	limit := limiter.MaxAllowPerSecond(limitValue)
	return func(ctx *gin.Context) {

		if !limit.AllowUri(ctx.Request.URL.Path) {
			response.Wrap(ctx).Error(http.StatusServiceUnavailable, "StatusServiceUnavailable")
			ctx.Abort()
		}
		ctx.Next()
	}
}
