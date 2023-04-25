package log

import (
	"github.com/kenSevLeb/go-framework/component/trace"
	"github.com/kenSevLeb/go-framework/util/net"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func init() {
	// 设置日志格式为json格式　自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	logrus.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	logrus.SetLevel(logrus.InfoLevel)
}

//var formatter = new(Formatter)
var _localHost = net.GetIP()

// Content 日志内容
type Content map[string]interface{}

func logger() *logrus.Entry {
	return logrus.WithField("trace_id", trace.Get()).WithField("host", _localHost)
}

// Info 信息
func Info(msg string, content Content) {
	logger().WithField("content", content).Info(msg)
}

// Error 错误
func Error(msg string, content Content) {
	logger().WithField("content", content).Error(msg)
}

// Debug 调试
func Debug(msg string, content Content) {
	logger().WithField("content", content).Debug(msg)
}

// SetOutput 指定日志输出
func SetOutput(out io.Writer) {
	logrus.SetOutput(out)
}
