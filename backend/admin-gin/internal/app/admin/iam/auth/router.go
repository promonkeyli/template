package auth

import (
	"mall-api/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup, handlers *Handler) {

	// 不需鉴权
	publicGroup := r.Group("/auth")
	{
		publicGroup.POST("/register", handlers.Register)
		publicGroup.POST("/login", handlers.Login)
		publicGroup.POST("/refresh", handlers.RefreshToken)
	}

	// 需要鉴权
	authGroup := r.Group("/auth")
	authGroup.Use(middleware.JWT())
	{
		// authGroup.POST("/logout", handlers.Logout)
	}
}
