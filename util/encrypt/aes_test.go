package encrypt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAesEncrypt(t *testing.T) {
	_, err := AesEncrypt("xx", "invalidKey")
	assert.NotNil(t, err)
	key := "hello_world20210"
	origin := fmt.Sprintf("aes:%d", time.Now().UnixNano())
	encrypt, err := AesEncrypt(origin, key)
	assert.Nil(t, err)
	fmt.Println(encrypt)
	decrypt, err := AesDecrypt(encrypt, key)
	assert.Nil(t, err)
	fmt.Println(decrypt)
	assert.Equal(t, origin, decrypt)
}

func TestAesDecrypt(t *testing.T) {
	key := "hello_world20210"
	_, err := AesDecrypt("wroneCrypt", key)
	assert.NotNil(t, err)
}
