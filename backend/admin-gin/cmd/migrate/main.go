// 数据库迁移
package main

import (
	"flag"
	"log/slog"
	"time"
)

func main() {
	var (
		dryRun  bool
		timeout time.Duration
	)

	flag.BoolVar(&dryRun, "dry-run", false, "只打印将要执行的迁移，不实际执行（当前仅做占位）")
	flag.DurationVar(&timeout, "timeout", 15*time.Second, "数据库连接与迁移超时时间")
	flag.Parse()

	slog.Info("开始执行数据库迁移...", "dryRun", dryRun, "timeout", timeout.String(), "configPath", "./configs")

	// 	pgCfg, err := database.ProvidePostgreConfig(cfg)
	// 	if err != nil {
	// 		slog.Error("解析数据库配置失败", "error", err.Error())
	// 		os.Exit(1)
	// 	}

	// 	db, err := database.NewPostgre(pgCfg)
	// 	if err != nil {
	// 		slog.Error("数据库连接失败", "error", err.Error())
	// 		os.Exit(1)
	// 	}

	// 	// 在这里注册需要迁移的模型
	// 	if err := db.AutoMigrate(
	// 		&user.User{},
	// 	); err != nil {
	// 		slog.Error("数据库迁移失败", "error", err.Error())
	// 		os.Exit(1)
	// 	}

	//		slog.Info("数据库迁移完成")
	//	}
}
