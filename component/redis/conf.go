package redis

import (
	"github.com/go-redis/redis"
	"net"
	"strconv"
	"time"
)

// redis配置
type Config struct {
	// 地址
	Host string

	// 端口
	Port int

	// 密码
	Auth string

	// db
	DB int

	// 连接池大小
	PoolSize int

	// 最大尝试次数
	MaxRetries int

	// 超时时间
	IdleTimeout time.Duration

	// 空闲链接检测频率
	IdleCheckFrequency time.Duration

	// 集群连接，逗号隔开
	Cluster []string

	// session的db
	SessionDB int
}

// 单个配置
func (conf *Config) Options() *redis.Options {
	address := net.JoinHostPort(conf.Host, strconv.Itoa(conf.Port))

	return &redis.Options{
		Addr:               address,
		Password:           conf.Auth,
		PoolSize:           conf.PoolSize,
		MaxRetries:         conf.MaxRetries,
		IdleTimeout:        conf.IdleTimeout,
		IdleCheckFrequency: conf.IdleCheckFrequency,
		DB:                 conf.DB,
	}
}

// 集群配置
func (conf *Config) ClusterOptions() *redis.ClusterOptions {
	return &redis.ClusterOptions{
		Addrs:              conf.Cluster,
		Password:           conf.Auth,
		PoolSize:           conf.PoolSize,
		MaxRetries:         conf.MaxRetries,
		IdleTimeout:        conf.IdleTimeout,
		IdleCheckFrequency: conf.IdleCheckFrequency,
	}
}

// 基于redis的session配置
func (conf *Config) SessionOptions() *redis.Options {
	address := net.JoinHostPort(conf.Host, strconv.Itoa(conf.Port))
	if conf.SessionDB == 0 {
		conf.SessionDB = conf.DB
	}

	return &redis.Options{
		Addr:        address,
		Password:    conf.Auth,
		PoolSize:    conf.PoolSize,
		MaxRetries:  conf.MaxRetries,
		IdleTimeout: conf.IdleTimeout,
		DB:          conf.SessionDB,
	}
}
