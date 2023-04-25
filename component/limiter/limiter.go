// 基于time/rate实现的限流器
// 封装，适合扩展使用
package limiter

import (
	cmap "github.com/orcaman/concurrent-map"
	"golang.org/x/time/rate"
	"time"
)

type Limiter struct {
	*rate.Limiter

	m cmap.ConcurrentMap
	v int
}

// 每秒允许最大访问量
func MaxAllowPerSecond(limitValue int) *Limiter {
	return &Limiter{v: limitValue, m: cmap.New(), Limiter: getRateLimit(limitValue)}
}

func getRateLimit(limitValue int) *rate.Limiter {
	r := rate.Every(time.Second / time.Duration(limitValue))
	limit := rate.NewLimiter(r, limitValue)
	return limit
}

// 根据uri判断
func (l *Limiter) AllowUri(uri string) bool {
	if limit, ok := l.m.Get(uri); ok {
		return limit.(*rate.Limiter).Allow()
	}

	limit := getRateLimit(l.v)
	l.m.Set(uri, limit)
	return limit.Allow()
}
