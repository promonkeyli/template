package router

import (
	adminAuth "mall-api/internal/app/admin/iam/auth"
	adminUser "mall-api/internal/app/admin/user"
	adminWire "mall-api/internal/app/wire"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func Router(r *gin.Engine, d *gorm.DB, rdb *redis.Client) {

	// openapi 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	handlers, err := adminWire.InitAdminHandlers(d, rdb)
	if err != nil {
		panic(err)
	}

	// admin 模块路由注册
	{
		// auth 路由注册
		adminAuth.RegisterRouter(r, handlers.Auth)

		// user 路由注册
		adminUser.RegisterRouter(r, handlers.User)
	}
}
