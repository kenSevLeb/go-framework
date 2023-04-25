package encrypt

import "encoding/base64"

// Base64Encode base64加密
func Base64Encode(key string) string {
	return base64.StdEncoding.EncodeToString([]byte(key))
}

// Base64Decode base64解密
func Base64Decode(encoded string) string {
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	return string(decoded)
}
