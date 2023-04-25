package http

type Config struct {
	// 运行环境
	RunMode string

	// 服务名称
	Name string

	// 服务端口
	Port int

	// http超时时间，单位:秒
	HttpRequestTimeOut int

	// 全局ID的header名称
	TraceHeader string

	// json web token signature
	JwtSign []byte

	// 是否开启调试日志
	LogDebug bool

	// 性能调试
	Pprof bool

	// 是否开启日志文档
	Swagger bool

	// 是否开启定时任务
	Task bool
}

const (
	EnvOnline = "online"
	EnvDev    = "dev"
	EnvLocal  = "local"
)

// 判断是否为线上环境
func (conf *Config) IsOnline() bool {
	return conf.RunMode == EnvOnline
}

// 判断是否为测试环境
func (conf *Config) IsDev() bool {
	return conf.RunMode == EnvDev
}

// 判断是否为本地环境
func (conf *Config) IsLocal() bool {
	return conf.RunMode == EnvLocal
}
