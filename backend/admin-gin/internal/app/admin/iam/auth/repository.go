package auth

import (
	"mall-api/internal/app/admin/user"

	"gorm.io/gorm"
)

type Repository interface {
	FindUserIsExist(name string) (bool, error)      // 查找用户是否存在
	CreateUser(user *user.User) error               // 创建新用户
	FindUserByName(name string) (*user.User, error) // 根据用户名查找用户
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

// true: 用户存在；false: 用户不存在
func (r *repo) FindUserIsExist(name string) (bool, error) {
	var count int64
	if err := r.db.Model(&user.User{}).Where("username = ?", name).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// 新增用户记录
func (r *repo) CreateUser(user *user.User) error {
	return r.db.Create(user).Error
}

// 查找用户记录
func (r *repo) FindUserByName(name string) (*user.User, error) {
	var user user.User

	err := r.db.Where("username = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
