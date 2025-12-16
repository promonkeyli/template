package database

import (
	"fmt"

	"mall-api/configs"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// ProviderSet provides config-driven constructors for PostgreSQL and Redis.
//
// Best practice used here:
// - Config is loaded in cmd/* via viper and passed into internal/wire.InitApp(cfg)
// - Providers in this package are pure (no filesystem/env reads)
// - We validate required fields early and return actionable errors
var ProviderSet = wire.NewSet(
	ProvidePostgreConfig,
	ProvideRedisConfig,
	NewPostgre,
	NewRedis,
)

// ProvidePostgreConfig maps *configs.Config into *database.PostgreConfig.
func ProvidePostgreConfig(cfg *configs.Config) (*PostgreConfig, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config is nil")
	}

	c := cfg.Database

	if c.Host == "" {
		return nil, fmt.Errorf("config.database.host is empty")
	}
	if c.User == "" {
		return nil, fmt.Errorf("config.database.user is empty")
	}
	if c.DBName == "" {
		return nil, fmt.Errorf("config.database.dbname is empty")
	}
	if c.Port == 0 {
		return nil, fmt.Errorf("config.database.port is zero")
	}
	if c.TimeZone == "" {
		return nil, fmt.Errorf("config.database.timezone is empty")
	}

	return &PostgreConfig{
		Host:     c.Host,
		User:     c.User,
		Password: c.Password,
		DBName:   c.DBName,
		Port:     c.Port,
		TimeZone: c.TimeZone,
	}, nil
}

// ProvideRedisConfig maps *configs.Config into *database.RedisConfig.
func ProvideRedisConfig(cfg *configs.Config) (*RedisConfig, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config is nil")
	}

	c := cfg.Redis

	if c.Addr == "" {
		return nil, fmt.Errorf("config.redis.addr is empty")
	}
	// Password can be empty.
	// DB can be 0.

	return &RedisConfig{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	}, nil
}

// Compile-time assertions (keep these so provider outputs stay aligned with Wire graph).
var (
	_ *gorm.DB
	_ *redis.Client
)
