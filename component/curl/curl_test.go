package curl

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var address = "http://localhost:8080/check"
var client = New(WithTimeout(time.Second * 30))

func TestGet(t *testing.T) {
	response, err := client.Get(address)
	assert.Nil(t, err)
	fmt.Println(response.String())
}

func TestPost(t *testing.T) {
	response, err := client.Post(address, nil)
	assert.Nil(t, err)
	fmt.Println(response.String())
}

func TestPut(t *testing.T) {
	response, err := client.Put(address, nil)
	assert.Nil(t, err)
	fmt.Println(response.String())
}

func TestPatch(t *testing.T) {
	response, err := client.Patch(address, nil)
	assert.Nil(t, err)
	fmt.Println(response.String())
}

func TestDelete(t *testing.T) {
	response, err := client.Delete(address, nil)
	assert.Nil(t, err)
	fmt.Println(response.String())
}

func TestResponse(t *testing.T) {

	request := client.NewRequest("GET", address, nil)
	if err := request.Err(); err != nil {
		t.Fatal(request.Err())

	}
	request.Header.Set("token", "123")
	response, err := request.Send()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(response.String())
	fmt.Println(string(response.Byte()))
	fmt.Println(response.Header())

	respObj := struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
	}{}
	assert.Nil(t, response.BindJson(&respObj))
	fmt.Println(respObj.Message)

}
