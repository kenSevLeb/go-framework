package app

import (
	"testing"
	"time"

	"go.uber.org/dig"
)

type comp1 struct{}

// Hey, dig is not intended to be invoked from your system's hot path. We expect
//it to be invoked at most once during startup, and definitely not concurrently.
//To discourage usage on the hot path, we have kept the APIs thread-unsafe.
func TestDI(t *testing.T) {
	di := dig.New()

	_ = di.Provide(func() *comp1 {
		return &comp1{}
	})

	for i := 0; i < 10000; i++ {
		go func() {
			_ = di.Invoke(func(c *comp1) {
				time.Sleep(1 * time.Millisecond)
			})
		}()
	}
}
