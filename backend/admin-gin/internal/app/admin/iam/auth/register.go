package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Register(rg *gin.RouterGroup, db *gorm.DB, rdb *redis.Client) {
	repo := NewRepository(db, rdb)
	svc := NewService(repo)
	h := NewHandler(svc)

	RegisterRouter(rg, h)
}
