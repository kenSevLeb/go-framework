package config

import (
	"errors"
	"git.yingxiong.com/platform/go-framework/component/log"
	"git.yingxiong.com/platform/go-framework/util/number"
	"git.yingxiong.com/platform/go-framework/util/strings"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

var (
	ErrEmptyFile = errors.New("empty file")
)

// 实例化
func New() *Config {
	conf := new(Config)
	conf.Viper = viper.New()
	return conf
}

// 通过文件加载配置
func (conf *Config) LoadFile(path ...string) error {
	if len(path) == 0 {
		return ErrEmptyFile
	}

	for _, item := range path {
		conf.SetConfigFile(item)
		if err := conf.MergeInConfig(); err != nil {
			return err
		}
	}
	if err := log.Init(lumberjackConfProvider(conf)); err != nil {
		return err
	}
	return nil
}

// 获取string型，并设置默认值
func (conf *Config) GetStringWithDefault(key string, dv string) string {
	return strings.Default(conf.GetString(key), dv)
}

// 获取int型，并设置默认值
func (conf *Config) GetIntWithDefault(key string, dv int) int {
	return number.DefaultInt(conf.GetInt(key), dv)
}
