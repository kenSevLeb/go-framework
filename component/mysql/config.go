package mysql

// MultiConfig 组合配置
type MultiConfig struct {
	items map[string]Config
}

// NewConfig 实例化MultiConfig
func NewConfig(items map[string]Config) MultiConfig {
	return MultiConfig{items: items}
}

// Get 获取配置项
func (mc MultiConfig) Get(name string) (Config, bool) {
	item, exist := mc.items[name]
	return item, exist
}

// Iterator 遍历配置项
func (mc MultiConfig) Iterator(fn func(name string, config Config)) {
	for name, config := range mc.items {
		fn(name, config)
	}
}

// Config 单个配置项
type Config struct {
	// 最大空闲连接数
	MaxIdleConnections int `mapstructure:"maxIdleConnections"`
	// 最大打开连接数
	MaxOpenConnections int `mapstructure:"maxOpenConnections"`
	// 最长活跃时间,单位：s
	MaxLifeTime int `mapstructure:"maxLifeTime"`
	// 空闲时间,单位：s
	MaxIdleTime int `mapstructure:"maxIdleTime"`
	// 主库连接
	Sources []string `json:"sources" mapstructure:"sources"`
	// 从库连接
	Replicas []string `json:"replicas" mapstructure:"replicas"`
	// 指定表
	Tables []string `json:"tables" mapstructure:"tables"` //
}
