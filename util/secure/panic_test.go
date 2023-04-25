package secure

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanic(t *testing.T) {
	var err error
	// 不会触发panic
	Panic(err)

	err = fmt.Errorf("test")
	// 会触发panic
	assert.Panics(t, func() {
		Panic(err)
	})

}
