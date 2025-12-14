package config

type ServerConfig struct {
	/** 主机ip地址 */
	Host string
	/** 数据库登录用户名 */
	User string
	/** 数据库登录密码 */
	Password string
	/** 数据库名称 */
	DBName string
	/** 数据库端口 */
	Port int
	/** 数据库时区 */
	TimeZone string
}

func NewServerConfig() ServerConfig {
	return ServerConfig{
		Host:     "",
		User:     "",
		Password: "",
		DBName:   "",
		Port:     0,
		TimeZone: "",
	}
}
