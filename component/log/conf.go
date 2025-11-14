package log

type LoggerMode string

const (
	ModeLumberjack LoggerMode = "lumberjack"
	ModeTimeRotate LoggerMode = "time"
)

type Config struct {
	Switch     bool       //开关
	Mode       LoggerMode // "lumberjack" 或 "time"
	LogPath    string     //日志目录 默认项目目录
	FileName   string     //日志名字 info.log
	MaxSize    int        //megabytes
	MaxBackups int        //文件数
	MaxAge     int        //days
	Compress   bool       //是否压缩
}
