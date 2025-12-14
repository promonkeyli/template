package router

import (
	adminAuth "mall-api/internal/app/admin/iam/auth"
	adminUser "mall-api/internal/app/admin/user"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Router(r *gin.Engine, d *gorm.DB, rdb *redis.Client) {

	// OpenAPI 路由注册
	RegisterOpenAPIRouter(r)

	// admin 模块路由注册
	{
		// auth 路由注册
		adminAuth.RegisterRouter(r, d, rdb)

		// user 路由注册
		adminUser.RegisterRouter(r, d, rdb)
	}
}
