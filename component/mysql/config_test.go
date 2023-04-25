package mysql

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiConfig_Get(t *testing.T) {
	multiConfig := NewConfig(map[string]Config{
		"default": {Sources: []string{"source1", "source2"}},
	})
	config, exist := multiConfig.Get("default")
	assert.True(t, exist)
	assert.NotNil(t, config)
}

func TestMultiConfig_Iterator(t *testing.T) {
	multiConfig := NewConfig(map[string]Config{
		"default": {Sources: []string{"source1", "source2"}},
		"other":   {Sources: []string{"otherSource"}},
	})

	multiConfig.Iterator(func(name string, config Config) {
		fmt.Printf("connect %s with source:%v\n", name, config.Sources)
	})
}

func TestNewConfig(t *testing.T) {
	multiConfig := NewConfig(map[string]Config{
		"default": {Sources: []string{"source1", "source2"}},
	})
	assert.NotNil(t, multiConfig)
}
