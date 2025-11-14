package task

import "git.yingxiong.com/platform/go-framework/util/strings"

// Option 任务选项接口
type Option interface {
	apply(*taskOption)
}

type taskOption struct {
	name string // 任务名称
}

// defaultOption 默认选项
func defaultOption() taskOption {
	return taskOption{name: strings.UUID()}
}

// 任务名称选项
type nameOption string

func (o nameOption) apply(opts *taskOption) {
	opts.name = string(o)
}

// WithName 设置任务名称
func WithName(name string) Option {
	return nameOption(name)
}
