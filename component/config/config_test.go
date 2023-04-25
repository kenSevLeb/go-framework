package config

import (
	"fmt"
	"kenSevLeb/go-framework/util/array"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	conf := New()
	assert.Nil(t, conf.LoadFile("/usr/app.yaml"))
	fmt.Println(conf.GetString("app.id"))

	fmt.Println(conf.GetStringSlice("app.id"), len(conf.GetStringSlice("app.id")))
	fmt.Println(conf.GetIntSlice("app.id"))
	fmt.Println(array.String2Int(conf.GetStringSlice("app.id")))
	fmt.Println(conf.GetString("redis.host"))
}

func TestConfig_LoadFile(t *testing.T) {
	conf := New()
	assert.Nil(t, conf.LoadFile("/usr/app.yaml", "/usr/other.yaml"))
	fmt.Println(conf.GetStringSlice("app.id"))
	fmt.Println(conf.GetString("test.name"))

}
