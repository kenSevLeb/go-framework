package kafka

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient(Config{
		Host: "10.0.6.190:9092",
		GroupId: "someGroup",
	})

	for {
		result, err := client.Read("test")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(result))
		}

		time.Sleep(time.Second)

	}

}

func TestClient_Write(t *testing.T) {
	client := NewClient(Config{
		Host: "10.0.6.190:9092",
		GroupId: "someGroup",
	})

	for i := 0; i < 100; i++ {

		err := client.Write("test", []byte(fmt.Sprintf("testContent:%d", i)))
		assert.Nil(t, err)
		time.Sleep(time.Second)
	}

}
