package log

import (
	"io"
	"os"
	"sync"

	"git.yingxiong.com/platform/go-framework/component/trace"
	"git.yingxiong.com/platform/go-framework/util/net"
	"github.com/sirupsen/logrus"
)

var (
	global *logger
	mu     sync.Mutex
)

type Writer interface {
	GetWriter(level logrus.Level, customPath ...string) io.WriteCloser
	Close() error
}

type logger struct {
	logger *logrus.Logger
	writer Writer
	conf   *Config
	mu     sync.Mutex
}

// Content 日志内容
type Content map[string]interface{}

func newClient(conf *Config) (*logger, error) {
	l := &logger{
		conf:   conf,
		logger: logrus.New(),
	}

	l.logger.SetFormatter(&logrus.JSONFormatter{})

	var writer Writer
	switch conf.Mode {
	case ModeTimeRotate:
		writer = NewTimeRotateLogger(conf)
	default:
		writer = NewLogger(conf)
	}

	l.writer = writer
	return l, nil
}

func (l *logger) Log(level logrus.Level, msg string, content Content, customPath ...string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	entry := l.logger.WithField("trace_id", trace.Get()).WithField("host", net.GetIP())
	if l.conf.Switch {
		w := l.writer.GetWriter(level, customPath...)
		l.logger.SetOutput(w)
	} else {
		l.logger.SetOutput(os.Stdout)
	}

	switch level {
	case logrus.InfoLevel:
		entry.WithField("content", content).Info(msg)
	case logrus.WarnLevel:
		entry.WithField("content", content).Warn(msg)
	case logrus.ErrorLevel:
		entry.WithField("content", content).Error(msg)
	case logrus.DebugLevel:
		entry.WithField("content", content).Debug(msg)
	}
}

func Init(conf *Config) (err error) {
	mu.Lock()
	defer mu.Unlock()
	if global == nil {
		global, err = newClient(conf)
	}
	return
}

func Close() {
	mu.Lock()
	defer mu.Unlock()
	if global != nil && global.writer != nil {
		_ = global.writer.Close()
	}
}

// Info 信息
func Info(msg string, content Content, customPath ...string) {
	if global != nil {
		global.Log(logrus.InfoLevel, msg, content, customPath...)
	}
}

// Error 错误
func Error(msg string, content Content, customPath ...string) {
	if global != nil {
		global.Log(logrus.ErrorLevel, msg, content, customPath...)
	}
}

// Debug 调试
func Debug(msg string, content Content, customPath ...string) {
	if global != nil {
		global.Log(logrus.DebugLevel, msg, content, customPath...)
	}
}
