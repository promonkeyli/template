package user

import (
	"context"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	// List 分页查询用户列表（仅返回未删除数据），支持 role 精确筛选与 keyword 模糊匹配（uid/username/email）
	List(ctx context.Context, page, size int, role, keyword string) ([]User, int, error)

	// Create 新增用户
	Create(ctx context.Context, u *User) error

	// GetByUID 按 UID 获取用户（用于更新前读取）
	GetByUID(ctx context.Context, uid string) (*User, error)

	// UpdateByUID 按 UID 更新（部分字段更新）
	UpdateByUID(ctx context.Context, uid string, updates map[string]any) error

	// SoftDeleteByUID 软删除（is_deleted=true）
	SoftDeleteByUID(ctx context.Context, uid string) error

	// ExistsByUsername 判断用户名是否存在（未删除）
	ExistsByUsername(ctx context.Context, username string) (bool, error)

	// ExistsByEmail 判断邮箱是否存在（未删除）
	ExistsByEmail(ctx context.Context, email string) (bool, error)

	// ExistsByEmailExcludeUID 判断邮箱是否存在（排除某个 UID，用于更新时校验）
	ExistsByEmailExcludeUID(ctx context.Context, email, excludeUID string) (bool, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) baseQuery(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Model(&User{}).Where("is_deleted = ?", false)
}

func (r *repo) List(ctx context.Context, page, size int, role, keyword string) ([]User, int, error) {
	q := r.baseQuery(ctx)

	role = strings.TrimSpace(role)
	if role != "" {
		q = q.Where("role = ?", role)
	}

	keyword = strings.TrimSpace(keyword)
	if keyword != "" {
		like := "%" + keyword + "%"
		q = q.Where("(uid ILIKE ? OR username ILIKE ? OR email ILIKE ?)", like, like, like)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var list []User
	if err := q.Order("created_at DESC").
		Offset((page - 1) * size).
		Limit(size).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, int(total), nil
}

func (r *repo) Create(ctx context.Context, u *User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *repo) GetByUID(ctx context.Context, uid string) (*User, error) {
	var u User
	if err := r.db.WithContext(ctx).Where("uid = ?", uid).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *repo) UpdateByUID(ctx context.Context, uid string, updates map[string]any) error {
	return r.db.WithContext(ctx).Model(&User{}).Where("uid = ?", uid).Updates(updates).Error
}

func (r *repo) SoftDeleteByUID(ctx context.Context, uid string) error {
	return r.db.WithContext(ctx).Model(&User{}).Where("uid = ?", uid).Updates(map[string]any{
		"is_deleted": true,
	}).Error
}

func (r *repo) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	var count int64
	if err := r.baseQuery(ctx).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repo) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return false, nil
	}

	var count int64
	if err := r.baseQuery(ctx).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repo) ExistsByEmailExcludeUID(ctx context.Context, email, excludeUID string) (bool, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return false, nil
	}

	var count int64
	if err := r.baseQuery(ctx).Where("email = ? AND uid <> ?", email, excludeUID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
