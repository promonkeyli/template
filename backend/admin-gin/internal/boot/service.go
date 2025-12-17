package boot

import (
	"mall-api/internal/app/admin/iam/auth"
	"mall-api/internal/app/admin/user"
)

type Services struct {
	Auth auth.Service
	User user.Service
}

func BuildServices(repos *Repos) *Services {
	return &Services{
		Auth: auth.NewService(repos.Auth),
		User: user.NewService(repos.User),
	}
}
