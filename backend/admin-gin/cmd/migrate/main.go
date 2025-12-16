// 数据库迁移
package main

import (
	"mall-api/pkg/logger"
	"os/user"
)

func main() {
	logger.Log.Info("开始执行数据库自动迁移...")

	// 在这里注册需要迁移的模型
	err := db.AutoMigrate(
		&user.User{},
	)

	if err != nil {
		logger.Log.Error("数据库迁移失败！")
	}

	logger.Log.Info("数据库自动迁移完成！")
}
