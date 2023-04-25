package curl

import "time"

// options 选项
type options struct {
	// 超时
	timeout time.Duration
	// traceId的header
	traceHeader string
}

type Option interface {
	apply(*options)
}

type timeoutOption time.Duration

func (t timeoutOption) apply(o *options) {
	o.timeout = time.Duration(t)
}

func WithTimeout(timeout time.Duration) Option {
	return timeoutOption(timeout)
}

type traceHeaderOption string

func (t traceHeaderOption) apply(o *options) {
	o.traceHeader = string(t)
}

func WithTraceHeader(traceHeader string) Option {
	return traceHeaderOption(traceHeader)
}
