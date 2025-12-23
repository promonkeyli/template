package user

import (
	"mall-api/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup, handlers *Handler) {
	ug := r.Group("/user")
	ug.Use(middleware.JWT())
	{
		ug.GET("", handlers.List)
		// ug.POST("", handlers.Create)
		// ug.PUT("/:id", handlers.Update)
		// ug.DELETE("/:id", handlers.Delete)
	}
}
