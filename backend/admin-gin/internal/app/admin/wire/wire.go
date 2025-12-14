//go:build wireinject
// +build wireinject

package wire

import (
	adminAuth "mall-api/internal/app/admin/iam/auth"
	adminUser "mall-api/internal/app/admin/user"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// AdminHandlers is the admin-module handler set resolved by Wire.
// Router registration should depend on these handlers rather than constructing them inside router funcs.
type AdminHandlers struct {
	Auth *adminAuth.Handler
	User *adminUser.Handler
}

// Provider sets for each module (repo -> service -> handler).
//
// Note:
//   - Wire can use concrete provider functions directly.
//   - `wire.Bind` is only needed when you provide a concrete type and want to satisfy an interface,
//     but here `NewService` takes the interface type already, so extra bindings are unnecessary.
var (
	authSet = wire.NewSet(
		adminAuth.NewRepository,
		adminAuth.NewService,
		adminAuth.NewHandler,
	)

	userSet = wire.NewSet(
		adminUser.NewRepository,
		adminUser.NewService,
		adminUser.NewHandler,
	)

	adminSet = wire.NewSet(
		authSet,
		userSet,
		wire.Struct(new(AdminHandlers), "Auth", "User"),
	)
)

// InitAdminHandlers wires up all admin handlers.
// Keep db/rdb provided by your bootstrap layer; wire only builds module graph.
func InitAdminHandlers(db *gorm.DB, rdb *redis.Client) (*AdminHandlers, error) {
	wire.Build(adminSet)
	return nil, nil
}
