package json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

func init() {
	extra.RegisterFuzzyDecoders()
}

// json序列化
func Encode(v interface{}) (string, error) {
	return jsoniter.MarshalToString(v)
}

// 忽略错误，只返回结果
func MustEncode(v interface{}) string {
	str, _ := Encode(v)
	return str
}

// json反序列化，obj为指针
func Decode(buf []byte, obj interface{}) error {
	return jsoniter.Unmarshal(buf, obj)
}

// 支持将json的string类型转为int/float等类型
func DecodeForceFromString(s string, obj interface{}) error {
	return jsoniter.UnmarshalFromString(s, &obj)
}

// DecodeForce
func DecodeForce(b []byte, obj interface{}) error {
	return jsoniter.Unmarshal(b, obj)
}
