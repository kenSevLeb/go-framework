package task

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithName(t *testing.T) {
	option := WithName("test")
	assert.Equal(t, nameOption("test"), option)
}
