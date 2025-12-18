package auth

import "time"

// account 是 auth 领域关心的最小账号信息（用于登录/注册）
type account struct {
	UID      string
	Username string
	Password string // bcrypt hash
	Role     string
	IsActive bool
}

// auth模块 user 仅用于 GORM 映射同一张用户表（与 user 模块解耦），仅模块内部使用
type user struct {
	ID uint64 `gorm:"primaryKey;autoIncrement"`

	UID      string `gorm:"size:32;uniqueIndex"`
	Username string `gorm:"size:64;uniqueIndex;not null"`
	Password string `gorm:"size:255;not null"`

	Role     string `gorm:"size:32;default:'admin'"`
	IsActive bool   `gorm:"default:true"`

	IsDeleted bool `gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName 强制 auth.user 使用 user 表
func (user) TableName() string { return "user" }
