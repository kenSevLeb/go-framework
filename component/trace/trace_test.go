package trace

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	New()
	assert.NotEmpty(t, Get())
}

func TestSet(t *testing.T) {
	expectTraceId := "someUniqueId"
	Set(expectTraceId)
	assert.Equal(t, expectTraceId, Get())
}

func TestGC(t *testing.T) {
	New()
	assert.NotEmpty(t, Get())
	GC()
	assert.Empty(t, Get())
}

func TestGo(t *testing.T) {
	New()
	outTraceId := Get()
	Go(func() {
		innerTraceId := Get()
		assert.Equal(t, outTraceId, innerTraceId)
	})
}
