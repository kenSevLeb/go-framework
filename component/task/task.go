package task

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/robfig/cron/v3"
)

// Task 定时任务接口
type Task interface {
	// AddJob 添加job
	AddJob(spec string, job cron.Job, opts ...Option) error

	// AddFunc 添加func
	AddFunc(spec string, job cron.FuncJob, opts ...Option) error

	// Has 判断是否已存在任务
	Has(name string) bool

	// Remove 删除任务
	Remove(name string)

	// Start 启动任务
	Start()

	// Stop 停止任务
	Stop()
}

// task 定时任务
type task struct {
	// instance 实例
	instance *cron.Cron

	// entryIds 任务ID映射,key为任务名称
	// 通过名称可以找到任务标识，从而实现删除任务的功能
	entryIds cmap.ConcurrentMap
}

// New 实例化
func New() Task {
	return &task{
		instance: cron.New(),
		entryIds: cmap.New(),
	}
}

// AddJob 添加任务
func (m *task) AddJob(spec string, job cron.Job, opts ...Option) error {
	options := defaultOption()
	for _, opt := range opts {
		opt.apply(&options)
	}

	if m.Has(options.name) {
		return nil
	}

	entryId, err := m.instance.AddJob(spec, job)
	if err != nil {
		return fmt.Errorf("AddJob:[%s],%v", options.name, err)
	}

	m.entryIds.Set(options.name, entryId)
	return nil
}

// AddFunc 添加任务
func (m *task) AddFunc(spec string, job cron.FuncJob, opts ...Option) error {
	options := defaultOption()
	for _, opt := range opts {
		opt.apply(&options)
	}

	if m.Has(options.name) {
		return nil
	}

	entryId, err := m.instance.AddFunc(spec, job)
	if err != nil {
		return fmt.Errorf("AddJob:[%s],%v", options.name, err)
	}

	m.entryIds.Set(options.name, entryId)
	return nil
}

// Has 判断任务是否已存在
func (m *task) Has(name string) bool {
	return m.entryIds.Has(name)
}

// Remove 删除任务
func (m *task) Remove(name string) {
	id, ok := m.entryIds.Get(name)
	if !ok {
		return
	}
	m.instance.Remove(id.(cron.EntryID))
	m.entryIds.Remove(name)
}

func (m *task) Start() {
	m.instance.Start()
}

func (m *task) Stop() {
	m.instance.Stop()
}
