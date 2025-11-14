package curl

import (
	"bytes"
	"errors"
	"git.yingxiong.com/platform/go-framework/util/conv"
	"git.yingxiong.com/platform/go-framework/util/json"
	"io"
	"net/http"
)

type Response interface {
	String() string
	Byte() []byte
	BindJson(object interface{}) error
	Reader() io.Reader
	StatusCode() int
}

// 响应
type response struct {
	// 状态码
	statusCode int
	// 响应头
	header http.Header
	// 采用字节存储
	body       []byte

}

var (
	ErrNoResponse = errors.New("no response")
)

// 返回字符串
func (resp *response) String() string {
	return conv.Byte2String(resp.body)
}

// 返回字节
func (resp *response) Byte() []byte {
	return resp.body
}

// 通过json解析到变量，object为指针类型
func (resp *response) BindJson(object interface{}) error {
	if resp.String() == "" {
		return ErrNoResponse
	}
	return json.Decode(resp.body, object)
}

// 返回reader
func (resp *response) Reader() io.Reader {
	return bytes.NewReader(resp.body)
}

// 返回状态码
func (resp *response) StatusCode() int {
	return resp.statusCode
}

// 返回header
func (resp *response) Header() http.Header {
	return resp.header
}
