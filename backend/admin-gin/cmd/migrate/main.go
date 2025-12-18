// 数据库迁移
package main

import (
	"log/slog"
	"mall-api/configs"
	"mall-api/internal/app/admin/user"
	"mall-api/internal/pkg/database"
	"os"
)

func main() {

	// 1.从环境变量获取运行模式
	runMode := os.Getenv("APP_ENV_MODE")
	env := "dev"
	if runMode != "" {
		env = runMode
	}

	// 2. 依据环境变量初始化系统配置
	cfg, err := configs.InitConfig(env)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// 3.连接数据库
	db, dbErr := database.NewPostgre(&database.PostgreConfig{
		Host:            cfg.Database.Host,
		User:            cfg.Database.User,
		Password:        cfg.Database.Password,
		DBName:          cfg.Database.DBName,
		Port:            cfg.Database.Port,
		TimeZone:        cfg.Database.TimeZone,
		MaxIdleConns:    cfg.Database.MaxIdleConns,
		MaxOpenConns:    cfg.Database.MaxOpenConns,
		ConnMaxLifetime: cfg.Database.ConnMaxLifetime,
	})
	if dbErr != nil {
		slog.Error(dbErr.Error())
		os.Exit(1)
	}

	// 4. 获取底层的 sql.DB 对象用于关闭
	sqlDB, err := db.DB()
	if err != nil {
		slog.Error(err.Error())
	}

	// 5. 使用 defer 确保在 main 函数结束时关闭连接
	defer func() {
		if err := sqlDB.Close(); err != nil {
			slog.Error("关闭数据库连接失败：%v", "error", err.Error())
		} else {
			slog.Info("数据库连接已关闭")
		}
	}()

	// 6. 迁移数据库
	if err := db.AutoMigrate(&user.User{}); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("数据库迁移成功")
}
