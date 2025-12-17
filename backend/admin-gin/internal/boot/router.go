package boot

import (
	"mall-api/internal/app/admin/iam/auth"
	"mall-api/internal/app/admin/user"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRouter(r *gin.Engine, h *Handlers) {

	// openapi 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 管理端路由
	adminGroup := r.Group("/admin")
	{
		auth.RegisterRouter(adminGroup, h.Auth)
		user.RegisterRouter(adminGroup, h.User)
	}
}
