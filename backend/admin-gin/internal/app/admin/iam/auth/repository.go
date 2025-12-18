package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Repository interface {
	FindUserIsExist(username string) (bool, error)    // 查找用户是否存在
	CreateUser(account *account) error                // 创建新用户
	FindUserByName(username string) (*account, error) // 根据用户名查找用户

	SetRefreshToken(ctx context.Context, uid string, token string, duration time.Duration) error // 设置刷新令牌
	GetRefreshToken(ctx context.Context, uid string) (string, error)                             // 设置刷新令牌
	DelRefreshToken(ctx context.Context, uid string) error                                       // 删除刷新令牌
}

type repo struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewRepository(db *gorm.DB, rdb *redis.Client) Repository {
	return &repo{db: db, rdb: rdb}
}

// 查找数据库是否存在该用户，true: 用户存在；false: 用户不存在
func (r *repo) FindUserIsExist(username string) (bool, error) {
	var count int64
	if err := r.db.Model(&user{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// 新增用户记录
func (r *repo) CreateUser(account *account) error {
	m := &user{
		UID:       account.UID,
		Username:  account.Username,
		Password:  account.Password,
		Role:      account.Role,
		IsActive:  account.IsActive,
		IsDeleted: false,
	}
	return r.db.Create(m).Error
}

// 查找用户记录
func (r *repo) FindUserByName(username string) (*account, error) {
	var m user
	if err := r.db.Where("username = ?", username).First(&m).Error; err != nil {
		return nil, err
	}

	return &account{
		UID:      m.UID,
		Username: m.Username,
		Password: m.Password,
		Role:     m.Role,
		IsActive: m.IsActive,
	}, nil
}

// 设置 refresh token
func (r *repo) SetRefreshToken(ctx context.Context, uid string, token string, duration time.Duration) error {
	// Key 格式建议: "模块:功能:ID" -> "auth:refresh:10086"
	key := fmt.Sprintf("auth:refresh:%s", uid)

	// Redis 命令: SET key value EX duration
	// duration (比如 7天) 会自动转换为 Redis 的秒数
	return r.rdb.Set(ctx, key, token, duration).Err()
}

// 获取 refresh token
func (r *repo) GetRefreshToken(ctx context.Context, uid string) (string, error) {
	key := fmt.Sprintf("auth:refresh:%s", uid)

	// Redis 命令: GET key
	val, err := r.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		// 返回自定义的通用错误，屏蔽底层细节
		return "", errors.New("Refresh Token Not Found")
	}

	// 处理其他错误（如 Redis 连接断开）
	if err != nil {
		return "", err
	}
	return val, nil
}

// 删除 refresh token
func (r *repo) DelRefreshToken(ctx context.Context, uid string) error {
	key := fmt.Sprintf("auth:refresh:%s", uid)
	return r.rdb.Del(ctx, key).Err()
}
