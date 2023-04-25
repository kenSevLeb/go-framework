package config

const (
	// 运行环境
	runModeKey = "server.runmode"

	// 服务名称配置项
	serverNameKey = "server.name"

	// 服务端口配置项
	serverPortKey = "server.port"

	// http请求超时配置项
	httpRequestTimeoutKey = "server.httpRequestTimeout"

	// 全局ID的header配置项
	traceHeaderKey = "server.traceHeader"

	// json web token
	jwtSignKey = "server.jwtSign"

	// 是否开启debug日志
	logDebugKey = "server.logDebug"

	// ip白名单
	ipFilterKey = "server.ipWhiteList"

	// 是否开启swagger
	swaggerSwitchKey = "server.swagger"

	// 是否开启pprof
	pprofSwitchKey = "server.pprof"

	// 是否开启task
	taskSwitchKey = "server.task"
)

const (
	redisHost               = "redis.host"
	redisDB                 = "redis.db"
	redisSessionDB          = "redis.sessionDB"
	redisPass               = "redis.pass"
	redisPort               = "redis.port"
	redisPoolSize           = "redis.poolSize"
	redisMaxRetries         = "redis.maxRetries"
	redisIdleTimeout        = "redis.idleTimeout"
	redisCluster            = "redis.cluster"
	redisIdleCheckFrequency = "redis.idleCheckFrequency"
)

const (
	esHost = "es.host"

	esIndexPrefix = "es.index_prefix"
)

const (
	dbKey = "db"
)

const (
	i18nFileKey = "i18n.files"
)
