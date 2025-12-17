// 应用组装入口
package boot

import (
	"log/slog"
	"mall-api/configs"
	"mall-api/internal/pkg/middleware"
	"mall-api/pkg/database"
	"mall-api/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	Log *slog.Logger
	Db  *gorm.DB
	Rdb *redis.Client
	Ge  *gin.Engine
}

func NewApp(cfg *configs.Config) (*App, error) {

	// 1. 构造 log
	log := logger.NewLog(logger.Config{
		Level:    cfg.Log.Level,
		Format:   cfg.Log.Format,
		Dir:      cfg.Log.Dir,
		Filename: cfg.Log.Filename,
		MaxSize:  cfg.Log.MaxSize,
		MaxAge:   cfg.Log.MaxAge,
		Compress: cfg.Log.Compress,
	})
	slog.SetDefault(log) // 设置全局默认slog，这里设置了，全局调用slog的地方都会用这个log

	// 2. 构造 gorm
	db, dbErr := database.NewPostgre(&database.PostgreConfig{
		Host:     cfg.Database.Host,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
		Port:     cfg.Database.Port,
		TimeZone: cfg.Database.TimeZone,
		// MaxIdleConns:    cfg.Database.MaxIdleConns,
		// MaxOpenConns:    cfg.Database.MaxOpenConns,
		// ConnMaxLifetime: cfg.Database.ConnMaxLifetime,
	})
	if dbErr != nil {
		return nil, dbErr
	}

	// 3. 构造 redis
	rdb, rdbErr := database.NewRedis(&database.RedisConfig{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	if rdbErr != nil {
		return nil, rdbErr
	}

	// 4. 构造 gin(使用干净的 Gin 引擎，方便接管日志以及其他中间件)
	ge := gin.New()
	ge.Use(
		gin.Recovery(),      // 1. 最外层兜底
		middleware.Cors(),   // 2. 尽早处理 OPTIONS
		middleware.Log(log), // 3. 正常请求日志
	)

	// 5. 构造 app
	app := &App{
		Log: log,
		Db:  db,
		Rdb: rdb,
		Ge:  ge,
	}
	return app, nil
}
