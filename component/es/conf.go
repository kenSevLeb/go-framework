package es

type Config struct {
	// es地址，如: 127.0.0.1:9200
	Host string

	// 索引前缀
	IndexPrefix string
	Username    string
	Pass        string
}
