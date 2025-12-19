package auth

import (
	"mall-api/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func registerRouter(r *gin.RouterGroup, handlers *handler) {

	// 不需鉴权
	publicGroup := r.Group("/auth")
	{
		publicGroup.POST("/register", handlers.register)
		publicGroup.POST("/login", handlers.login)
		publicGroup.POST("/refresh", handlers.refresh)
	}

	// 需要鉴权
	authGroup := r.Group("/auth")
	authGroup.Use(middleware.JWT())
	{
		authGroup.POST("/logout", handlers.logout)
	}
}
