//go:build wireinject
// +build wireinject

package wire

import (
	"mall-api/configs"
	adminWire "mall-api/internal/app/admin/wire"
	"mall-api/internal/router"
	"mall-api/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// App is the fully-wired application object returned by the injector.
type App struct {
	Engine *gin.Engine
	DB     *gorm.DB
	RDB    *redis.Client

	// Keep these as fields so Wire considers the providers "used" and also so
	// the application graph guarantees these side-effects happen during init.
	AdminHandlers     *adminWire.AdminHandlers
	RoutersRegistered routersRegistered
}

// routersRegistered is a no-op marker type used to model "route registration happened"
// as a Wire-provided dependency.
type routersRegistered struct{}

// provideGinEngine constructs a bare gin engine. Middlewares are configured in main
// in this repo; if you want to move middleware wiring into Wire, do it here.
func provideGinEngine() *gin.Engine {
	return gin.New()
}

// provideAdminHandlers wires up admin module handlers (auth/user/etc).
func provideAdminHandlers(db *gorm.DB, rdb *redis.Client) (*adminWire.AdminHandlers, error) {
	return adminWire.InitAdminHandlers(db, rdb)
}

// provideRouters registers routes onto gin.Engine (side-effect) and returns a marker.
// Wire providers must return at least one output, so we return a no-op type here.
func provideRouters(r *gin.Engine, db *gorm.DB, rdb *redis.Client) routersRegistered {
	router.Router(r, db, rdb)
	return routersRegistered{}
}

// InitApp builds the full application graph: DB, Redis, Gin Engine, module handlers,
// and routes registration.
//
// Best practice: config is loaded in cmd/* (via viper) and passed in explicitly,
// so the injector has no filesystem/env side-effects.
func InitApp(cfg *configs.Config) (*App, error) {
	wire.Build(
		// Input config
		wire.Value(cfg),

		// Infra (Postgre/Redis built from viper-managed config)
		database.ProvidePostgreConfig,
		database.ProvideRedisConfig,
		database.NewPostgre,
		database.NewRedis,

		// Gin engine
		provideGinEngine,

		// Admin modules (forces construction early; also validates the graph)
		provideAdminHandlers,

		// Routes (side-effect registration)
		provideRouters,

		// Final app aggregate (include AdminHandlers and RoutersRegistered so providers are used)
		wire.Struct(new(App), "Engine", "DB", "RDB", "AdminHandlers", "RoutersRegistered"),
	)
	return nil, nil
}
