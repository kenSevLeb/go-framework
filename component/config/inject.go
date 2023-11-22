package config

import (
	"github.com/kenSevLeb/go-framework/component/es"
	"github.com/kenSevLeb/go-framework/component/i18n"
	"github.com/kenSevLeb/go-framework/component/kafka"
	"github.com/kenSevLeb/go-framework/component/log"
	"github.com/kenSevLeb/go-framework/component/mongo"
	"github.com/kenSevLeb/go-framework/component/mysql"
	"github.com/kenSevLeb/go-framework/component/redis"
	"github.com/kenSevLeb/go-framework/http"
	"github.com/kenSevLeb/go-framework/rpc"
	"go.uber.org/dig"
	"time"
)

func Inject(container *dig.Container) error {
	_ = container.Provide(New)
	_ = container.Provide(httpConfigProvider)
	if err := container.Provide(mysqlConfigProvider); err != nil {
		return err
	}
	_ = container.Provide(redisConfigProvider)
	_ = container.Provide(elasticSearchConfigProvider)
	_ = container.Provide(mongoConfigProvider)
	_ = container.Provide(i18nConfigProvider)
	_ = container.Provide(rpcConfigProvider)
	_ = container.Provide(kafkaConfigProvider)
	_ = container.Provide(lumberjackConfProvider)
	return nil
}

func httpConfigProvider(conf *Config) *http.Config {
	return &http.Config{
		RunMode:            conf.GetStringWithDefault(runModeKey, "local"),
		Name:               conf.GetStringWithDefault(serverNameKey, "app"),
		Port:               conf.GetIntWithDefault(serverPortKey, 8080),
		HttpRequestTimeOut: conf.GetIntWithDefault(httpRequestTimeoutKey, 10),
		TraceHeader:        conf.GetStringWithDefault(traceHeaderKey, "request-trace"),
		LogDebug:           conf.GetBool(logDebugKey),
		JwtSign:            []byte(conf.GetString(jwtSignKey)),
		Swagger:            conf.GetBool(swaggerSwitchKey),
		Pprof:              conf.GetBool(pprofSwitchKey),
		Task:               conf.GetBool(taskSwitchKey),
	}
}

func mysqlConfigProvider(conf *Config) (mysql.MultiConfig, error) {
	items := make(map[string]mysql.Config)
	var mc mysql.MultiConfig
	if err := conf.UnmarshalKey(dbKey, &items); err != nil {
		return mc, err
	}
	mc = mysql.NewConfig(items)
	return mc, nil
}

func redisConfigProvider(conf *Config) *redis.Config {
	return &redis.Config{
		Host:               conf.GetString(redisHost),
		Port:               conf.GetInt(redisPort),
		Auth:               conf.GetString(redisPass),
		DB:                 conf.GetInt(redisDB),
		PoolSize:           conf.GetIntWithDefault(redisPoolSize, 50),
		MaxRetries:         conf.GetIntWithDefault(redisMaxRetries, 3),
		IdleTimeout:        time.Second * time.Duration(conf.GetIntWithDefault(redisIdleTimeout, 3)),
		IdleCheckFrequency: time.Second * time.Duration(conf.GetIntWithDefault(redisIdleCheckFrequency, 60)),
		Cluster:            conf.GetStringSlice(redisCluster),
		SessionDB:          conf.GetInt(redisSessionDB),
	}
}

func elasticSearchConfigProvider(conf *Config) *es.Config {
	return &es.Config{
		Host:        conf.GetString(esHost),
		IndexPrefix: conf.GetString(esIndexPrefix),
	}
}

func mongoConfigProvider(conf *Config) *mongo.Config {
	return &mongo.Config{
		Hosts:    conf.GetStringSlice("mongo.host"),
		Source:   conf.GetString("mongo.source"),
		Username: conf.GetString("mongo.username"),
		Password: conf.GetString("mongo.password"),
		Timeout:  conf.GetIntWithDefault("mongo.timeout", 60),
		Database: conf.GetString("mongo.database"),
	}
}

func i18nConfigProvider(conf *Config) *i18n.Config {
	return &i18n.Config{
		Files: conf.GetStringSlice(i18nFileKey),
	}
}

func rpcConfigProvider(conf *Config) *rpc.Config {
	return &rpc.Config{
		Port: conf.GetIntWithDefault("rpc.port", 9000),
	}
}

func kafkaConfigProvider(conf *Config) *kafka.Config {
	return &kafka.Config{
		Host:    conf.GetString("kafka.host"),
		GroupId: conf.GetString("kafka.groupId"),
	}
}

func lumberjackConfProvider(conf *Config) *log.Config {
	return &log.Config{
		Switch:     conf.GetBool("lumberjack.switch"),
		LogPath:    conf.GetString("lumberjack.logPath"),
		FileName:   conf.GetString("lumberjack.fileName"),
		MaxSize:    conf.GetInt("lumberjack.maxSize"),
		MaxBackups: conf.GetInt("lumberjack.maxBackups"),
		MaxAge:     conf.GetInt("lumberjack.maxAge"),
		Compress:   conf.GetBool("lumberjack.compress"),
	}
}
