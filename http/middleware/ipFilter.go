package middleware

import (
	"kenSevLeb/go-framework/util/set"
	"github.com/gin-gonic/gin"
)

// IPFilter 白名单过滤器
func IPFilter(ips []string) gin.HandlerFunc {

	s := set.ThreadSafe()
	for _, ip := range ips {
		s.Add(ip)
	}
	return func(ctx *gin.Context) {
		if !s.Empty() && !s.Contain(ctx.ClientIP()) {
			ctx.AbortWithStatus(403)
			return
		}

		ctx.Next()
	}
}
