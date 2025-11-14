package log

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

type TimeRotateLogger struct {
	mu              sync.Mutex
	loggers         map[string]*lumberjack.Logger
	basePath        string
	fileName        string
	dateFormat      string
	conf            *Config
	lastCleanupTime time.Time
}

func NewTimeRotateLogger(conf *Config) *TimeRotateLogger {
	return &TimeRotateLogger{
		loggers:    make(map[string]*lumberjack.Logger),
		basePath:   conf.LogPath,
		fileName:   conf.FileName,
		dateFormat: "2006-01-02",
		conf:       conf,
	}
}

func (t *TimeRotateLogger) Write(p []byte) (n int, err error) {
	writer := t.getCurrentWriter()
	return writer.Write(p)
}

func (t *TimeRotateLogger) getCurrentWriter() io.Writer {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.maybeCleanup()

	return t
}

func extractDateFromFilename(filename string) (string, bool) {
	base := filepath.Base(filename)
	if len(base) >= 10 {
		return base[:10], true
	}
	return "", false
}

func (t *TimeRotateLogger) cleanupOldLoggers() {
	today := time.Now().Format(t.dateFormat)
	for key, log := range t.loggers {
		if date, ok := extractDateFromFilename(log.Filename); ok && date != today {
			_ = log.Close()
			delete(t.loggers, key)
		}
	}
	t.lastCleanupTime = time.Now()
}

func (t *TimeRotateLogger) maybeCleanup() {
	// 每次10%的概率触发，或者距离上次10分钟清除
	now := time.Now()
	if now.Sub(t.lastCleanupTime) > 10*time.Minute && rand.Intn(100) > 90 {
		t.cleanupOldLoggers()
	}
}

func (t *TimeRotateLogger) Close() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	for _, log := range t.loggers {
		_ = log.Close()
	}
	t.loggers = nil
	return nil
}

func (t *TimeRotateLogger) GetWriter(level logrus.Level, customPath ...string) io.WriteCloser {
	//key := level.String()
	today := time.Now().Format(t.dateFormat)

	key := "log" // 感觉不同日志等级都在一个文件中好些，可以直接看到上下文，而不用去各个文件中找
	if len(customPath) > 0 {
		key = filepath.Base(customPath[0])
	}
	mapKey := fmt.Sprintf("%s-%s", key, today)

	t.mu.Lock()
	defer t.mu.Unlock()

	// 清除不用的key，避免map无限膨胀
	t.maybeCleanup()

	if t.loggers == nil {
		t.loggers = make(map[string]*lumberjack.Logger)
	}
	if log, exists := t.loggers[mapKey]; exists {
		return log
	}

	logPath, err := os.Getwd()
	if err != nil {
		logPath = "." // 兜底，避免panic
	}
	fullPath := fmt.Sprintf("%s/log/%s/%s.log", logPath, key, today)
	_ = os.MkdirAll(filepath.Dir(fullPath), 0755)

	log := &lumberjack.Logger{
		Filename:   fullPath,
		MaxSize:    t.conf.MaxSize,
		MaxBackups: t.conf.MaxBackups,
		MaxAge:     t.conf.MaxAge,
		Compress:   t.conf.Compress,
	}
	t.loggers[mapKey] = log
	return log
}
