package limiter

import (
	"fmt"
	"testing"
)

func BenchmarkMaxAllowPerSecond(b *testing.B) {
	limit := MaxAllowPerSecond(1000)
	for i := 0; i < b.N; i++ {
		if limit.Allow() {
			fmt.Println("allow")
		} else {
			fmt.Println("limit")
		}
	}
}
