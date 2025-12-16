// 数据库迁移
package main

import (
	"flag"
	"os"
	"time"

	"mall-api/configs"
	"mall-api/pkg/database"
	"mall-api/pkg/logger"

	// 这里导入的是你真实的业务模型包（当前文件原来导入的是 os/user，是错误的）
	"mall-api/internal/app/admin/user"
)

func main() {
	// 初始化 slog 日志（保证 logger.Log 可用）
	logger.Init()

	var (
		dryRun  bool
		timeout time.Duration
	)

	flag.BoolVar(&dryRun, "dry-run", false, "只打印将要执行的迁移，不实际执行（当前仅做占位）")
	flag.DurationVar(&timeout, "timeout", 15*time.Second, "数据库连接与迁移超时时间")
	flag.Parse()

	// 读取 config.yaml（viper），固定路径：./configs
	cfg, err := configs.LoadConfig("./configs")
	if err != nil {
		logger.Log.Error("读取配置失败", "error", err)
		os.Exit(1)
	}

	logger.Log.Info("开始执行数据库迁移...", "dryRun", dryRun, "timeout", timeout.String(), "configPath", "./configs")

	pgCfg, err := database.ProvidePostgreConfig(cfg)
	if err != nil {
		logger.Log.Error("解析数据库配置失败", "error", err)
		os.Exit(1)
	}

	db, err := database.NewPostgre(pgCfg)
	if err != nil {
		logger.Log.Error("数据库连接失败", "error", err)
		os.Exit(1)
	}

	// 在这里注册需要迁移的模型
	if err := db.AutoMigrate(
		&user.User{},
	); err != nil {
		logger.Log.Error("数据库迁移失败", "error", err)
		os.Exit(1)
	}

	logger.Log.Info("数据库迁移完成")
}
