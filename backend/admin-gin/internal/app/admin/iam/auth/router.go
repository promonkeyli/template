package auth

import (
	"mall-api/internal/pkg/mw"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RegisterRouter(r *gin.Engine, d *gorm.DB, rdb *redis.Client) {

	// 数据层：负责数据库交互
	repository := NewRepository(d)

	// 业务层：负责核心业务逻辑
	service := NewService(repository)

	// 接入层：负责处理 HTTP 请求
	handler := NewHandler(service)

	// 不需要鉴权的路由
	publicGroup := r.Group("/admin/auth")
	{
		publicGroup.POST("/register", handler.Register)
		publicGroup.POST("/login", handler.Login)
	}

	// 需要鉴权的路由
	authGroup := r.Group("/admin/auth")
	authGroup.Use(mw.JWT())
	{
		// authGroup.POST("/logout", handler.Logout)
	}
}
