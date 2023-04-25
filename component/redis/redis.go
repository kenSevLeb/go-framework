package redis

import (
	"github.com/go-redis/redis"
	"log"
)

// redis客户端
type Client struct {
	redis.UniversalClient

	conf *Config
}

type ClusterClient struct {
	redis.UniversalClient
}

func newClient(conf *Config) (*Client, error) {
	connection := redis.NewClient(conf.Options())
	_, err := connection.Ping().Result()
	if err != nil {
		return nil, err
	}
	log.Println("connect redis success:", conf.Host)

	return &Client{UniversalClient: connection}, nil
}

func newClusterClient(conf *Config) (*ClusterClient, error) {
	connection := redis.NewClusterClient(conf.ClusterOptions())
	_, err := connection.Ping().Result()
	if err != nil {
		return nil, err
	}
	log.Println("connect redis cluster success:", conf.Cluster)

	return &ClusterClient{UniversalClient: connection}, nil
}
