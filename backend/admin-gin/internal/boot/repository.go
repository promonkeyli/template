package boot

import (
	"mall-api/internal/app/admin/iam/auth"
	"mall-api/internal/app/admin/user"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Repos struct {
	Auth auth.Repository
	User user.Repository
}

func BuildRepos(db *gorm.DB, rdb *redis.Client) *Repos {
	return &Repos{
		Auth: auth.NewRepository(db, rdb),
		User: user.NewRepository(db),
	}
}
