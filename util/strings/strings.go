package strings

import (
	bytes2 "bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// return the encrypt string by md5 algorithm
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// return hash string
func Hash(s string) string {
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(s))
	result := Sha1Inst.Sum([]byte(""))
	return fmt.Sprintf("%x", result)
}

// 返回唯一ID
func UUID() string {
	v4 := uuid.NewV4()
	return v4.String()
}

// 根据连接符号连接数组
func Implode(items interface{}, separator string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(items), "[]"), " ", separator, -1)
}

// 分割字符词，返回字符词数组
func Explode(str, separator string) []string {
	return strings.Split(str, separator)
}

// 返回默认值
func Default(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}

	return value
}

// 截取字符串
func SubStr(s string, l int) string {
	if len(s) <= l {
		return s
	}
	ss, sl, rl, rs := "", 0, 0, []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			rl = 1
		} else {
			rl = 2
		}

		if sl+rl > l {
			break
		}
		sl += rl
		ss += string(r)
	}
	return ss
}

var longLetters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// 生成随机字符串
// n 长度
func Rand(n int) string {
	rand.Seed(time.Now().UnixNano())
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return ""
	}
	for i, x := range b {
		arc = x & 61
		b[i] = longLetters[arc]
	}
	return string(b)
}

func GetRequestBody(req *http.Request) string {
	switch req.Method {
	case http.MethodGet:
		return req.URL.Query().Encode()

	case http.MethodPost:
		fallthrough
	case http.MethodPut:
		fallthrough
	case http.MethodPatch:
		var bodyBytes []byte // 我们需要的body内容

		bodyBytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return ""
		}
		req.Body = ioutil.NopCloser(bytes2.NewBuffer(bodyBytes))

		return string(bodyBytes)

	}

	return ""
}
