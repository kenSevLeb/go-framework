package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"kenSevLeb/go-framework/component/event"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"log"
	"time"
)

var (
	ErrUnknownDefaultConnection = errors.New("unknown default connection")
	ErrUnknownDsn               = errors.New("unknown dsn")
)

const (
	defaultConnection = "default"
)

// BaseEntity 基础Entity
type BaseEntity struct {
	// 主键
	Id int `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	// 软删除标识
	IsDel int `json:"-"`
	// 创建时间，格式为10位的时间戳
	CreatedAt int `json:"created_at"`
	// 更新时间，格式为10位的时间戳
	UpdatedAt int `json:"updated_at"`
}

// Columns 字段别名，可用于Update方法
type Columns map[string]interface{}

// Client mysql客户端
type Client struct {
	*gorm.DB

	configs MultiConfig
}

func (client *Client) GetInstance() *gorm.DB {
	return client.DB
}

// getResolverConf 解析配置
func getResolverConf(sourceItems, replicaItems []string) dbresolver.Config {
	var sources, replicas []gorm.Dialector
	for _, source := range sourceItems {
		sources = append(sources, mysql.Open(source))
	}

	for _, replica := range replicaItems {
		replicas = append(replicas, mysql.Open(replica))
	}

	return dbresolver.Config{
		Sources:  sources,                   // 主库配置
		Replicas: replicas,                  // 从库配置
		Policy:   dbresolver.RandomPolicy{}, // sources/replicas 负载均衡策略
	}
}

// 连接默认库
func (client *Client) connectDefault() error {
	config, ok := client.configs.Get(defaultConnection)
	if !ok {
		return ErrUnknownDefaultConnection
	}

	if len(config.Sources) == 0 {
		return ErrUnknownDsn
	}

	sqlDB, err := sql.Open("mysql", config.Sources[0])
	if err != nil {
		return err
	}

	// set pool config
	sqlDB.SetMaxIdleConns(config.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(config.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(config.MaxLifeTime))

	conn, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return fmt.Errorf("default database connect failed:%v", err)
	}
	client.DB = conn

	return nil
}

// 连接数据库
func (client *Client) connect() error {
	if err := client.connectDefault(); err != nil {
		return err
	}

	plugin := new(dbresolver.DBResolver)

	// 遍历
	client.configs.Iterator(func(name string, config Config) {
		tables := make([]interface{}, 0)
		for _, t := range config.Tables {
			tables = append(tables, t)
		}
		tables = append(tables, name)
		if name == defaultConnection { // 主连接已经初始化了connection
			plugin.Register(getResolverConf(nil, config.Replicas), tables...)
		} else {
			plugin.Register(getResolverConf(config.Sources, config.Replicas), tables...).
				SetMaxIdleConns(config.MaxIdleConnections).
				SetMaxOpenConns(config.MaxOpenConnections).
				SetConnMaxLifetime(time.Second * time.Duration(config.MaxLifeTime))
		}
	})

	if err := client.DB.Use(plugin); err != nil {
		return fmt.Errorf("use plugin: %v", err)
	}

	_ = event.Trigger(event.AfterDatabaseConnect, nil)

	log.Println("connect database success")

	return nil
}

// GetConnection 根据名称获取指定数据库连接
func (client *Client) GetConnection(name string) *gorm.DB {
	return client.DB.Clauses(dbresolver.Use(name))
}
