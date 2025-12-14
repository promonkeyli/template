package auth

import "time"

// Account 是 auth 领域关心的最小账号信息（用于登录/注册）。
// 它不应该泄露 user 模块的全量用户模型，从而实现 auth 与 user 解耦。
type Account struct {
	UID      string
	Username string
	Password string // bcrypt hash
	Role     string
	IsActive bool
}

// userModel 仅用于 GORM 映射同一张用户表（与 user 模块解耦）。
//
// 注意：如果不显式指定 TableName()，GORM 会把结构体名 userModel 映射成 user_model，
// 这会导致查询落到不存在的表。
// 这里强制它使用已有的 user 表。
type userModel struct {
	ID uint64 `gorm:"primaryKey;autoIncrement"`

	UID      string `gorm:"size:32;uniqueIndex"`
	Username string `gorm:"size:64;uniqueIndex;not null"`
	Password string `gorm:"size:255;not null"`

	Role     string `gorm:"size:32;default:'admin'"`
	IsActive bool   `gorm:"default:true"`

	// 为了与现有表结构对齐（你 user 模型里有该字段），保留它避免插入/查询时出现字段差异带来的困惑。
	IsDeleted bool `gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName 强制 auth.userModel 使用 user 表（避免默认映射成 user_model）
func (userModel) TableName() string { return "user" }
