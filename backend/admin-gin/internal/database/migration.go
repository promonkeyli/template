package database

import (
	"mall-api/internal/app/admin/user"
	"mall-api/internal/logger"

	"gorm.io/gorm"
)

// AutoMigrate 自动迁移数据库结构
func InitAutoMigrate(db *gorm.DB) error {
	logger.Log.Info("开始执行数据库自动迁移...")

	// 在这里注册需要迁移的模型
	err := db.AutoMigrate(
		&user.User{},
	)

	if err != nil {
		return err
	}

	logger.Log.Info("数据库自动迁移完成！")
	return nil
}
