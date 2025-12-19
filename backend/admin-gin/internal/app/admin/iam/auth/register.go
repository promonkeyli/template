package auth

import (
	"mall-api/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Register(rg *gin.RouterGroup, db *gorm.DB, rdb *redis.Client, jt *jwt.JWT) {
	repo := newRepository(db, rdb)
	svc := newService(repo, jt)
	h := newHandler(svc)

	registerRouter(rg, h)
}
