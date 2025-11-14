package log

import (
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

type Lumberjack struct {
	Logger   *lumberjack.Logger
	Switch   bool
	FileName string
}

func NewLogger(conf *Config) *Lumberjack {
	return &Lumberjack{
		Switch:   conf.Switch,
		FileName: conf.FileName,
		Logger: &lumberjack.Logger{
			MaxSize:    conf.MaxSize,
			MaxBackups: conf.MaxBackups,
			MaxAge:     conf.MaxAge,
			Compress:   conf.Compress,
		},
	}
}

func (l *Lumberjack) SetFilePath(level string) {
	logPath, _ := os.Getwd()
	logPath = logPath + "/log"
	l.Logger.Filename = logPath + "/" + level + "/" + l.FileName
}

func (l *Lumberjack) GetWriter(level logrus.Level, customPath ...string) io.WriteCloser {
	if len(customPath) > 0 {
		l.SetFilePath(customPath[0])
	} else {
		l.SetFilePath(level.String())
	}
	return l.Logger
}

func (l *Lumberjack) Close() error {
	if closer, ok := any(l.Logger).(io.Closer); ok {
		return closer.Close()
	}
	return nil
}
