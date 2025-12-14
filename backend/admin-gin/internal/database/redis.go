package database

import (
	"context"
	"mall-api/internal/config"
	"mall-api/internal/logger"

	"github.com/redis/go-redis/v9"
)

// InitRedis 初始化Redis连接
func InitRedis() (*redis.Client, error) {
	cfg := config.NewRedisConfig()

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// 测试连接
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	logger.Log.Info("Redis 连接成功！")
	return rdb, nil
}
