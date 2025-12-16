//	@title			Mall API
//	@version		1.0
//	@description	Mall API 服务接口文档
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8081
//	@BasePath	/

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.

package main

import (
	"os"

	"mall-api/configs"
	"mall-api/internal/pkg/middleware"
	"mall-api/internal/wire"
	"mall-api/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化 slog 日志
	logger.Init()

	// 读取配置（viper），固定从 ./configs/config.yaml 加载
	cfg, err := configs.LoadConfig("./configs")
	if err != nil {
		logger.Log.Error("读取配置失败", "error", err)
		os.Exit(1)
	}

	// 设置 gin mode：支持 debug / release / test（与 gin 内置一致）
	if cfg.Server.Mode != "" {
		switch cfg.Server.Mode {
		case gin.DebugMode, gin.ReleaseMode, gin.TestMode:
			gin.SetMode(cfg.Server.Mode)
		default:
			logger.Log.Error("非法的 server.mode（仅支持 debug/release/test）", "mode", cfg.Server.Mode)
			os.Exit(1)
		}
	}

	// 使用 Wire 初始化应用（显式传入配置）
	app, err := wire.InitApp(cfg)
	if err != nil {
		logger.Log.Error("应用初始化失败", "error", err)
		os.Exit(1)
	}

	// 数据库迁移已迁移到 cmd/migrate，不在服务启动时自动执行

	// 修改 Gin 所有日志使用 slog
	gin.DefaultWriter = logger.Writer()
	gin.DefaultErrorWriter = logger.Writer()

	// 使用 Wire 创建的 Engine（内部已完成路由注册）
	r := app.Engine

	// 移除该警告：[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
	r.SetTrustedProxies(nil)

	// 挂载日志中间件
	r.Use(middleware.SlogLog())

	// 挂载错误恢复中间件
	r.Use(middleware.SlogRecovery())

	// 挂载跨域中间件
	r.Use(middleware.Cors())

	// 启动服务
	r.Run(":8081")
}
