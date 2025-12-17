// viper 集中管理配置
package configs

import (
	"errors"
	"strings"

	"github.com/spf13/viper"
)

// Config 聚合所有配置
type Config struct {
	App      App      `mapstructure:"app"`
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
	Redis    Redis    `mapstructure:"redis"`
	JWT      JWT      `mapstructure:"jwt"`
	Log      Log      `mapstructure:"log"`
	CORS     CORS     `mapstructure:"cors"`
}

// App 应用基础配置
type App struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

// Server HTTP服务配置
type Server struct {
	Port         int    `mapstructure:"port"`
	Mode         string `mapstructure:"mode"`          // debug, release, test
	ReadTimeout  int    `mapstructure:"read_timeout"`  // 秒
	WriteTimeout int    `mapstructure:"write_timeout"` // 秒
}

// Database 数据库配置 (PostgreSQL)
type Database struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	DBName          string `mapstructure:"dbname"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	SSLMode         string `mapstructure:"ssl_mode"`
	TimeZone        string `mapstructure:"timezone"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"` // 秒
}

// Redis 缓存配置
type Redis struct {
	Addr         string `mapstructure:"addr"`
	Password     string `mapstructure:"password"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	DialTimeout  int    `mapstructure:"dial_timeout"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

// JWT 认证配置 (双Token)
type JWT struct {
	Secret        string `mapstructure:"secret"`
	Issuer        string `mapstructure:"issuer"`
	AccessExpire  int64  `mapstructure:"access_expire"`  // 秒
	RefreshExpire int64  `mapstructure:"refresh_expire"` // 秒
}

// Log 日志配置
type Log struct {
	Level    string `mapstructure:"level"`    // debug, info, warn, error
	Format   string `mapstructure:"format"`   // json, text
	Dir      string `mapstructure:"dir"`      // 日志文件夹
	Filename string `mapstructure:"filename"` // 日志文件名
	MaxSize  int    `mapstructure:"max_size"` // MB
	MaxAge   int    `mapstructure:"max_age"`  // 天
	Compress bool   `mapstructure:"compress"` // 是否压缩
}

// CORS 跨域配置
type CORS struct {
	AllowOrigins []string `mapstructure:"allow_origins"`
}

// ================================ viper 配置项目初始化 ===================================

func InitConfig(env string) (*Config, error) {
	// 1. 创建 viper 实例
	v := viper.New()

	// 2. 环境变量通用设置：db.host -> DB_HOST -> APP_DB_HOST
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 3. 开发环境从 yaml 文件中加载配置
	if env == "dev" {
		// 开发环境读取yaml文件配置
		v.SetConfigName("config") // 配置文件名 (不带扩展名)
		v.SetConfigType("yaml")   // 如果配置文件中没有扩展名，则明确指定
		v.AddConfigPath("./configs")
		if err := v.ReadInConfig(); err != nil {
			return nil, errors.New("读取配置文件失败")
		}
	}

	// 映射配置文件 然后返回
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, errors.New("配置文件配置有误")
	}
	return &cfg, nil
}
