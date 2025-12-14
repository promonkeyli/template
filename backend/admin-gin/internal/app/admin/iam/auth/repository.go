package auth

import "gorm.io/gorm"

type Repository interface {
	FindUserIsExist(username string) (bool, error)    // 查找用户是否存在
	CreateUser(account *Account) error                // 创建新用户
	FindUserByName(username string) (*Account, error) // 根据用户名查找用户
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

// true: 用户存在；false: 用户不存在
func (r *repo) FindUserIsExist(username string) (bool, error) {
	var count int64
	if err := r.db.Model(&userModel{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// 新增用户记录
func (r *repo) CreateUser(account *Account) error {
	m := &userModel{
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
func (r *repo) FindUserByName(username string) (*Account, error) {
	var m userModel
	if err := r.db.Where("username = ?", username).First(&m).Error; err != nil {
		return nil, err
	}

	return &Account{
		UID:      m.UID,
		Username: m.Username,
		Password: m.Password,
		Role:     m.Role,
		IsActive: m.IsActive,
	}, nil
}
