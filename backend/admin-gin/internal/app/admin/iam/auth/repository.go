package auth

import (
	"mall-api/internal/app/admin/staff"

	"gorm.io/gorm"
)

type Repository interface {
	FindStaffIsExist(name string) (bool, error)        // 查找用户是否存在
	CreateStaff(user *staff.Staff) error               // 创建新用户
	FindStaffByName(name string) (*staff.Staff, error) // 根据用户名查找用户
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

// true: 用户存在；false: 用户不存在
func (r *repo) FindStaffIsExist(name string) (bool, error) {
	var count int64
	if err := r.db.Model(&staff.Staff{}).Where("username = ?", name).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// 新增 staff 记录
func (r *repo) CreateStaff(staff *staff.Staff) error {
	return r.db.Create(staff).Error
}

// 查找 staff 记录
func (r *repo) FindStaffByName(name string) (*staff.Staff, error) {
	var staff staff.Staff

	err := r.db.Where("username = ?", name).First(&staff).Error
	if err != nil {
		return nil, err
	}

	return &staff, nil
}
