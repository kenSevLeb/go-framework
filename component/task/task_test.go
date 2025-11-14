package task

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	instance := New()
	assert.NotNil(t, instance)
}

func TestTask_AddFunc(t *testing.T) {
	instance := New()
	err := instance.AddFunc("@every 5s", func() {
		fmt.Println("some task")
	})
	assert.Nil(t, err)
}

type testJob struct{}

func (job testJob) Run() {
	fmt.Println("run")
}

func TestTask_AddJob(t *testing.T) {
	instance := New()
	err := instance.AddJob("@every 5s", new(testJob))
	assert.Nil(t, err)
}

func TestTask_Has(t *testing.T) {
	instance := New()
	jobName := "testTask"
	err := instance.AddFunc("@every 5s", func() {
		fmt.Println("some task")
	}, WithName(jobName))
	assert.Nil(t, err)
	assert.True(t, instance.Has(jobName))
}

func TestTask_Remove(t *testing.T) {
	instance := New()
	jobName := "testTask"
	err := instance.AddFunc("@every 5s", func() {
		fmt.Println("some task")
	}, WithName(jobName))
	assert.Nil(t, err)
	assert.True(t, instance.Has(jobName))

	instance.Remove(jobName)
	assert.False(t, instance.Has(jobName))
}

func TestTask_Start(t *testing.T) {
	New().Start()
}

func TestTask_Stop(t *testing.T) {
	instance := New()
	instance.Stop()
}
