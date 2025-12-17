package boot

import (
	"mall-api/internal/app/admin/iam/auth"
	"mall-api/internal/app/admin/user"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func Register(r *gin.Engine, db *gorm.DB, rdb *redis.Client) {
	// openapi routes
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// admin routes
	adminGroup := r.Group("/admin")
	{
		auth.Register(adminGroup, db, rdb)
		user.Register(adminGroup, db)
	}
}
