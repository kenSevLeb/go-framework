package log

import (
	"github.com/kenSevLeb/go-framework/component/trace"
	"github.com/kenSevLeb/go-framework/util/net"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	global, _ = newClient(&Config{
		Switch: false,
	})
)

type Logger interface {
	Log(level logrus.Level, msg string, content Content)
}

type logger struct {
	entry *logrus.Entry
	conf  *Config
}

// Content 日志内容
type Content map[string]interface{}

func newClient(conf *Config) (Logger, error) {
	// 设置日志格式为json格式　自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	return &logger{
		entry: logrus.WithField("trace_id", trace.Get()).WithField("host", net.GetIP()),
		conf:  conf,
	}, nil
}

func (logger *logger) Log(level logrus.Level, msg string, content Content) {
	outLog, isLumber := logger.conf.LumberjackOptions(level.String())
	if !isLumber {
		logrus.SetOutput(os.Stdout)
	} else {
		logrus.SetOutput(outLog)
	}
	switch level {
	case logrus.InfoLevel:
		logger.entry.WithField("content", content).Info(msg)
	case logrus.WarnLevel:
		logger.entry.WithField("content", content).Warn(msg)
	case logrus.ErrorLevel:
		logger.entry.WithField("content", content).Error(msg)
	case logrus.DebugLevel:
		logger.entry.WithField("content", content).Debug(msg)
	}
	if isLumber {
		//关闭文件
		_ = outLog.Close()
	}
}

func Init(cfg *Config) (err error) {
	global, err = newClient(cfg)
	return
}

// Info 信息
func Info(msg string, content Content) {
	global.Log(logrus.InfoLevel, msg, content)
}

// Error 错误
func Error(msg string, content Content) {
	global.Log(logrus.ErrorLevel, msg, content)
}

// Debug 调试
func Debug(msg string, content Content) {
	global.Log(logrus.DebugLevel, msg, content)
}
