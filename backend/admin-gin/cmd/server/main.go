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
	"mall-api/internal/database"
	"mall-api/internal/logger"
	"mall-api/internal/pkg/mw"
	"mall-api/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化 slog 日志
	logger.Init()

	// 初始化数据库连接
	db, err := database.InitDB()
	if err != nil {
		logger.Log.Error("数据库初始化失败", "error", err)
		return
	}

	// 数据库自动迁移
	if err := database.InitAutoMigrate(db); err != nil {
		logger.Log.Error("数据库迁移失败", "error", err)
		return
	}

	// 修改 Gin 所有日志使用 slog
	gin.DefaultWriter = logger.Writer()
	gin.DefaultErrorWriter = logger.Writer()

	// 移除该警告使用gin.New(): [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
	r := gin.New()

	// 移除该警告：[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
	r.SetTrustedProxies(nil)

	// 挂载日志中间件
	r.Use(mw.SlogLog())

	// 挂载错误恢复中间件
	r.Use(mw.SlogRecovery())

	// 挂载跨域中间件
	r.Use(mw.Cors())

	// JWT鉴权中间件将在具体的路由组中按需注册

	// 初始化 Redis 连接
	rdb, err := database.InitRedis()
	if err != nil {
		logger.Log.Error("Redis 初始化失败", "error", err)
		return
	}

	// 注册路由 (传入 handler)
	router.Router(r, db, rdb)

	// 启动服务
	r.Run(":8081")
}
