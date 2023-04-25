package mongo

type Config struct {
	// 地址
	Hosts []string

	//
	Source string

	// 用户名
	Username string

	// 密码
	Password string

	// 超时时间
	Timeout int

	// 数据库
	Database string
}
