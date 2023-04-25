package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// 加密
func AesEncrypt(orig string, key string) (string, error) {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = pkcS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	crypt := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(crypt, origData)
	return base64.StdEncoding.EncodeToString(crypt), nil
}

// 解密
func AesDecrypt(crypt string, key string) (string, error) {
	// 转成字节数组
	cryptByte, err := base64.StdEncoding.DecodeString(crypt)
	if err != nil {
		return "", err
	}
	k := []byte(key)
	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(cryptByte))
	// 解密
	blockMode.CryptBlocks(orig, cryptByte)
	// 去补全码
	orig = pkcS7UnPadding(orig)
	return string(orig), nil
}

// 补码
// AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func pkcS7Padding(cipherStr []byte, blocksize int) []byte {
	padding := blocksize - len(cipherStr)%blocksize
	PadContent := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherStr, PadContent...)
}

// 去码
func pkcS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
