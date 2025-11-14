package middleware

import (
	"net"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kenSevLeb/go-framework/util/set"
)

// IPFilter 白名单过滤器
func IPFilter(ips []string) gin.HandlerFunc {

	ipSet := set.ThreadSafe()
	var cidrs []*net.IPNet
	for _, v := range ips {
		if strings.Contains(v, "/") {
			if _, ipnet, err := net.ParseCIDR(v); err == nil {
				cidrs = append(cidrs, ipnet)
			}
			continue
		}
		ipSet.Add(v)
	}
	return func(ctx *gin.Context) {
		if ipSet.Empty() && len(cidrs) == 0 {
			ctx.Next()
			return
		}

		clientIP := ctx.ClientIP()
		if ipSet.Contain(clientIP) {
			ctx.Next()
			return
		}

		ip := net.ParseIP(clientIP)
		if ip != nil {
			for _, n := range cidrs {
				if n.Contains(ip) {
					ctx.Next()
					return
				}
			}
		}

		ctx.AbortWithStatus(403)
	}
}
