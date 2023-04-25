package lumberjack

import "github.com/natefinch/lumberjack"

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

// 集群配置
func (conf *Config) LumberjackOptions(logFileName string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    conf.MaxSize,
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge,
		Compress:   conf.Compress,
	}
}
