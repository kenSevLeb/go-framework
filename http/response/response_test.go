package response

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewResponse(t *testing.T) {
	response := NewResponse()
	assert.Equal(t, 0, response.Code)
	assert.Equal(t, "", response.Message)
}

func BenchmarkRequestIdWithPlus(b *testing.B) {
	id := uuid.NewV4().String()
	for i := 0; i < b.N; i++ {
		_ = requestPrefix + id
	}
}

func BenchmarkRequestIdWithStringWriter(b *testing.B) {
	id := uuid.NewV4().String()
	for i := 0; i < b.N; i++ {
		var result strings.Builder
		result.WriteString(requestPrefix)
		result.WriteString(id)
	}
}
