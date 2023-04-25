package trace

// trace 基于协程的服务全局链路追踪组件
// 一般在http的中间件，定时任务里初始化,业务里仅仅需要GET即可
// 使得同一个协程下的所有业务都使用一个ID标记，用于追踪业务执行链路
// 可用于日志分析，用户请求溯源
// Usage:
// 		trace.New()
//  	defer trace.GC() // 注：使用defer延迟调用，表示最后执行id的回收释放容量，否则容易导致内存溢出
//		fmt.Println(trace.Id())

import (
	"fmt"
	"github.com/kenSevLeb/go-framework/util/strings"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/petermattis/goid"
)

var (
	// 保存协程的全局唯一ID数组
	traceIds = cmap.New()
)

// New 快捷生成traceId的方法
func New() {
	Set(getUuid())
}

// Get 获取当前协程的唯一ID，如果没有设置当前协程的唯一ID，会返回空字符串
func Get() string {
	result, _ := traceIds.Get(goroutineId())
	if result == nil {
		return ""
	}
	return result.(string)
}

// GC 回收唯一ID,释放map容量
func GC() {
	traceIds.Remove(goroutineId())
}

// goroutineId 获取协程ID
func goroutineId() string {
	return fmt.Sprintf("g:%d", goid.Get())
}

// Go 在协程间传递traceId
func Go(f func()) {
	go func(id string) {
		Set(id)
		defer GC()
		f()
	}(Get())
}

func getUuid() string {
	return "trace:" + strings.UUID()
}

func Set(id string) {
	traceIds.Set(goroutineId(), id)
}
