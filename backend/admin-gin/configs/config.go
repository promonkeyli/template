package configs

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	Port     int    `mapstructure:"port"`
	TimeZone string `mapstructure:"timezone"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

var (
	once     sync.Once
	global   *Config
	globalEr error
)

func LoadConfig(path string) (*Config, error) {
	v := viper.New()

	// --- 设置读取路径和文件名 ---
	v.AddConfigPath(path)
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	// --- 环境变量覆盖 ---
	// 例如：MALL_DATABASE_HOST=127.0.0.1 覆盖 database.host
	v.SetEnvPrefix("MALL")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// --- 读取配置 ---
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("找不到配置文件 config.yml/config.yaml，请检查路径: %s", path)
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

// MustLoadGlobal 初始化全局配置（只会执行一次）。
// 你可以在程序启动时调用它，然后通过 Get() 获取配置。
func MustLoadGlobal(path string) {
	once.Do(func() {
		global, globalEr = LoadConfig(path)
	})
	if globalEr != nil {
		panic(globalEr)
	}
}

// Get 返回全局配置；在调用前请确保已执行 MustLoadGlobal。
func Get() *Config {
	return global
}

// Err 返回全局配置初始化错误（若有）。
func Err() error {
	return globalEr
}
