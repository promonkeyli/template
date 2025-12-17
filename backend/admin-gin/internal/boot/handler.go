package boot

import (
	"mall-api/internal/app/admin/iam/auth"
	"mall-api/internal/app/admin/user"
)

type Handlers struct {
	Auth *auth.Handler
	User *user.Handler
}

func BuildHandlers(services *Services) *Handlers {
	return &Handlers{
		Auth: auth.NewHandler(services.Auth),
		User: user.NewHandler(services.User),
	}
}
