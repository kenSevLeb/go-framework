package middleware

import (
	"git.yingxiong.com/platform/go-framework/component/trace"
	"github.com/gin-gonic/gin"
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
