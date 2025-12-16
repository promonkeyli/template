package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

// path: 配置文件所在的文件夹路径，例如 "./configs"
func LoadConfig(path string) (*Config, error) {
	v := viper.New()

	// --- 设置读取路径和文件名 ---
	v.AddConfigPath(path)     // 设置配置文件的搜索目录
	v.SetConfigName("config") // 设置文件名 (不带后缀，它会自动找 config.yaml, config.json 等)
	v.SetConfigType("yaml")   // 明确指定文件类型 (如果文件名没有后缀，这行是必须的)

	// --- (可选) 开启环境变量覆盖 ---
	// 允许通过环境变量 SERVER_PORT=9090 来覆盖 yaml 里的 server.port
	v.SetEnvPrefix("MALL")
	v.AutomaticEnv()

	// --- 读取配置 ---
	if err := v.ReadInConfig(); err != nil {
		// 判断是否是“找不到文件”的错误
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("找不到配置文件，请检查路径: %s", path)
		}
		return nil, fmt.Errorf("读取配置文件出错: %v", err)
	}

	// --- 映射到结构体 ---
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("配置解析失败(Unmarshal): %v", err)
	}

	return &c, nil
}
