package database

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Addr     string
	Password string
	DB       int

	// 连接池
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
}

func NewRedis(c *RedisConfig) (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Password:     c.Password,
		DB:           c.DB,
		DialTimeout:  time.Duration(c.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(c.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(c.WriteTimeout) * time.Second,
	})

	// 测试连接
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
