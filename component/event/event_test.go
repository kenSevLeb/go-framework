package event

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	eventType = "hello"
)

func TestEvent(t *testing.T) {

	Register(eventType, Listener{
		Mode: Sync,
		Handle: func(ev Event) {
			fmt.Println(ev.Params)
		},
	})

	err := Trigger(eventType, "test")
	assert.Nil(t, err)

	assert.True(t, Has(eventType))
	Remove(eventType)
	assert.False(t, Has(eventType))

}
