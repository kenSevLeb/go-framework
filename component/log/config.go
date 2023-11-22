package log

import (
	"fmt"
	"github.com/kenSevLeb/go-framework/util/strings"
	"github.com/natefinch/lumberjack"
	"os"
)

type Config struct {
	//Filename   string //日志目录
	Switch     bool   //开关
	LogPath    string //日志目录 默认项目目录
	FileName   string //日志名字 info.log
	MaxSize    int    //megabytes
	MaxBackups int    //文件数
	MaxAge     int    //days
	Compress   bool   //是否压缩
}

// 配置
func (conf *Config) LumberjackOptions(level string) (*lumberjack.Logger, bool) {
	if !conf.Switch {
		return nil, false
	}
	logPath := conf.LogPath
	if logPath == "" {
		logPath, _ = os.Getwd()
		logPath = fmt.Sprintf("%s/log", logPath)
	}
	fileName := fmt.Sprintf("%s/%s/%s", logPath, level, strings.Default(conf.FileName, "main.log"))
	return &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    conf.MaxSize,
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge,
		Compress:   conf.Compress,
	}, true
}
