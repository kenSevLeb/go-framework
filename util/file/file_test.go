package file

import (
	"fmt"
	"github.com/kenSevLeb/go-framework/util/date"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDir(t *testing.T) {
	fmt.Println(Dir("/root/test.go"))
}

func TestMkdir(t *testing.T) {
	err := Mkdir("/tmp/a/b/c", true)
	assert.Nil(t, err)
}

func TestExist(t *testing.T) {
	assert.True(t, Exist("/tmp/a/b/c"))
}

func TestDownloadRemoteFile(t *testing.T) {
	filename := fmt.Sprintf("%s/%s", "/tmp", date.GetLocalMicroTimeStampStr())
	assert.Nil(t, DownloadRemoteFile(filename, "http://assets.processon.com/chart_image/5f4869b9e0b34d1abc6c43f8.png"))
}

func TestGetCurrentDir(t *testing.T) {
	fmt.Println(GetCurrentDir())
}
