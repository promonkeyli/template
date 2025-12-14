package config

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// GetRedisConfig 获取Redis配置
// TODO: 后续可以改为从配置文件读取
func NewRedisConfig() RedisConfig {
	return RedisConfig{
		Addr:     "139.155.143.190:6379", // 假设与PG在同一服务器，或者使用 localhost:6379
		Password: "ly15984093508",        // 假设密码与DB用户相同，或者为空，这里先预留
		DB:       0,
	}
}
