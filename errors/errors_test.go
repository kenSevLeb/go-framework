package errors

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNew(t *testing.T) {
	code := 1001
	message := "invalid param"
	err := New(code, message)
	assert.Equal(t, err.Code, code)
	assert.Equal(t, err.Message, message)
}

func TestWith(t *testing.T) {
	err := Unauthorized("unauthorized").With("SomePrefix")
	fmt.Println(err.Error())
}

func TestError_Error(t *testing.T) {
	err := InternalServer("some error")
	fmt.Println(err.Error())
	fmt.Printf("%v\n", err)
}

func TestError_Hide(t *testing.T) {
	err := New(1001, "参数错误")
	fmt.Println(err.Error())
	hideErr := err.Hide("xxx").Append("附加信息")
	fmt.Println(hideErr.Error())
	fmt.Println(hideErr.ErrorWithHide())
	fmt.Println(New(1002, "查询数据失败").With("结果异常").Append("获取日志信息错误").Error())
}
