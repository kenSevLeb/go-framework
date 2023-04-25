package lumberjack

import (
	"fmt"
	"kenSevLeb/go-framework/component/trace"
	"kenSevLeb/go-framework/util/net"
	"kenSevLeb/go-framework/util/strings"
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	*logrus.Entry
	conf *Config
}

func newClient(conf *Config) (*Logger, error) {
	// 设置日志格式为json格式　自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 设置日志级别为warn以上
	logrus.SetLevel(logrus.InfoLevel)
	entity := logrus.
		WithField("trace_id", trace.Get()).
		WithField("host", net.GetIP())
	return &Logger{
		Entry: entity,
		conf:  conf,
	}, nil
}

type Content map[string]interface{}

func (impl Logger) Info(msg string, content Content) {
	impl.setPath("info")
	impl.Logger.WithField("content", content).Info(msg)
}

func (impl Logger) Error(msg string, content Content) {
	impl.setPath("error")
	impl.Logger.WithField("content", content).Error(msg)
}

func (impl Logger) Debug(msg string, content Content) {
	impl.setPath("debug")
	impl.Logger.WithField("content", content).Debug(msg)
}

func (impl Logger) setPath(level string) {
	if !impl.conf.Switch {
		impl.Logger.SetOutput(os.Stdout)
	} else {
		logPath := impl.conf.LogPath
		if logPath == "" {
			logPath, _ = os.Getwd()
			logPath = fmt.Sprintf("%s/log", logPath)
		}
		fileName := strings.Default(impl.conf.FileName, "project.log")
		fileName = fmt.Sprintf("%s/%s/%s", logPath, level, fileName)
		logger := impl.conf.LumberjackOptions(fileName)
		impl.Logger.SetOutput(logger)
	}
}
