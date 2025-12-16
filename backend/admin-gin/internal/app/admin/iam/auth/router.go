package auth

import (
	"mall-api/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine, handler *Handler) {

	// 不需要鉴权的路由
	publicGroup := r.Group("/admin/auth")
	{
		publicGroup.POST("/register", handler.Register)
		publicGroup.POST("/login", handler.Login)
		publicGroup.POST("/refresh", handler.RefreshToken)
	}

	// 需要鉴权的路由
	authGroup := r.Group("/admin/auth")
	authGroup.Use(middleware.JWT())
	{
		// authGroup.POST("/logout", handler.Logout)
	}
}
