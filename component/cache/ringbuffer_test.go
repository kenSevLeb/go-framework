package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRingQueue(t *testing.T) {
	rq := NewRingQueue(16)
	assert.NotNil(t, rq)
}

func TestRingQueue_IsEmpty(t *testing.T) {
	rq := NewRingQueue(16)
	assert.True(t, rq.IsEmpty())

	rq.Write(1)
	assert.False(t, rq.IsEmpty())

	_, _ = rq.Read()
	assert.True(t, rq.IsEmpty())
}

func TestRingQueue_Write(t *testing.T) {
	rq := NewRingQueue(16)
	rq.Write(1)
}

func TestRingQueue_Read(t *testing.T) {
	rq := NewRingQueue(16)
	rq.Write(1)
	item, err := rq.Read()
	assert.Nil(t, err)
	assert.Equal(t, 1, item)
}

func BenchmarkRingQueue_Read(b *testing.B) {
	rq := NewRingQueue(16)
	rq.Write(1)

	for i := 0; i < b.N; i++ {
		_, _ = rq.Read()
	}
}

func BenchmarkRingQueue_Write(b *testing.B) {
	rq := NewRingQueue(16)
	for i := 0; i < b.N; i++ {
		rq.Write(1)
	}
}

func TestRingQueue_Len(t *testing.T) {
	rq := NewRingQueue(16)
	assert.Equal(t, 0, rq.Len())

	rq.Write(1)
	assert.Equal(t, 1, rq.Len())

	_, _ = rq.Read()
	assert.Equal(t, 0, rq.Len())

	rq.Write(2)
	assert.Equal(t, 1, rq.Len())
}
