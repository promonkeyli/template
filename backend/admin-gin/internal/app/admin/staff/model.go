package staff

import "time"

type Staff struct {
	/** 自增主键（数据库内部使用，不对外暴露） */
	ID uint64 `gorm:"primaryKey;autoIncrement"`

	/** 全局唯一用户标识（对外使用）用于安全、日志、审计等，不暴露真实数据库 ID */
	UID string `gorm:"size:32;uniqueIndex"`

	/** 登录用户名 */
	Username string `gorm:"size:64;uniqueIndex;not null"`

	/** 邮箱（可选）可用于找回密码、内部通知等 */
	Email string `gorm:"size:128;uniqueIndex"`

	/** 密码哈希值（bcrypt 加密）绝不存储明文密码 */
	Password string `gorm:"size:255;not null"`

	/** 角色标识（RBAC 权限） 示例：admin / operator / finance / customer_service */
	Role string `gorm:"size:32;default:'staff'"`

	/** 是否启用账号（控制能否登录）*/
	IsActive bool `gorm:"default:true"`

	/** 是否软删除 */
	IsDeleted bool `gorm:"default:false"`

	/** 创建时间 */
	CreatedAt time.Time

	/** 更新时间 */
	UpdatedAt time.Time

	/** 注销时间 */
	DeleteAt time.Time
}
