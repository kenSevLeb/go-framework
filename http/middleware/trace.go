package middleware

import (
	"github.com/gin-gonic/gin"
	"kenSevLeb/go-framework/component/trace"
	"strings"
)

// Trace 全局链路中间件
func Trace(traceHeader string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 如果是其他服务的http请求，则从头部信息获取，接续链路
		traceId := strings.TrimSpace(ctx.GetHeader(traceHeader))
		if traceId == "" {
			trace.New()
		} else {
			trace.Set(traceId)
		}

		ctx.Next()

		trace.GC()
	}

}
