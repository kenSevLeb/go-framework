package rpc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServer(t *testing.T) {
	server := NewServer(&Config{Port: 8080})
	assert.Equal(t, 8080, server.conf.Port)
	assert.NotNil(t, server.Server)
}

func TestServer_Start(t *testing.T) {
	server := NewServer(&Config{Port: 5002})
	err := server.Start()
	assert.Nil(t, err)
}
