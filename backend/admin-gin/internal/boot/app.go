// 应用组装入口
package boot

import (
	"fmt"
	"log/slog"
	"mall-api/configs"
	"mall-api/internal/pkg/database"
	"mall-api/internal/pkg/jwt"
	"mall-api/internal/pkg/logger"
	"mall-api/internal/pkg/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	Log *slog.Logger
	Db  *gorm.DB
	Rdb *redis.Client
	Jt  *jwt.JWT
	Ge  *gin.Engine
	Se  *http.Server
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
		Host:            cfg.Database.Host,
		User:            cfg.Database.User,
		Password:        cfg.Database.Password,
		DBName:          cfg.Database.DBName,
		Port:            cfg.Database.Port,
		TimeZone:        cfg.Database.TimeZone,
		MaxIdleConns:    cfg.Database.MaxIdleConns,
		MaxOpenConns:    cfg.Database.MaxOpenConns,
		ConnMaxLifetime: cfg.Database.ConnMaxLifetime,
	})
	if dbErr != nil {
		return nil, dbErr
	}

	// 3. 构造 redis
	rdb, rdbErr := database.NewRedis(&database.RedisConfig{
		Addr:         cfg.Redis.Addr,
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
		DialTimeout:  cfg.Redis.DialTimeout,
		ReadTimeout:  cfg.Redis.ReadTimeout,
		WriteTimeout: cfg.Redis.WriteTimeout,
	})
	if rdbErr != nil {
		return nil, rdbErr
	}

	// 4. 构造 JWT 引擎，并初始化 jwt 中间件
	jt := jwt.New(
		cfg.JWT.Secret,
		cfg.JWT.Issuer,
		time.Duration(cfg.JWT.AccessExpire)*time.Second,
		time.Duration(cfg.JWT.RefreshExpire)*time.Second,
	)
	middleware.InitJWT(jt) // jwt 中间件注入jwt引擎，避免每次调用都传入

	// 5. 接管 Gin 内部日志、路由加载信息，全部转为 slog 形式（gin构造必须采用gin.New,且必须在gin.New()之前进行接管）
	logger.BuilderGinLog(log)

	// 6. 构造 gin(使用干净的 Gin 引擎，方便接管日志以及其他中间件)
	ge := gin.New()
	ge.Use(
		gin.Recovery(),    // 1. 最外层兜底
		middleware.Cors(), // 2. 尽早处理 OPTIONS
		middleware.Log(),  // 3. 正常请求日志
	)

	// 7. 构造 http.Server
	se := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      ge,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,  // 读取请求体的最大请求时间
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second, // 写入响应的最大请求时间
	}

	// 8. 构造 app
	app := &App{
		Log: log,
		Db:  db,
		Rdb: rdb,
		Jt:  jt,
		Ge:  ge,
		Se:  se,
	}
	return app, nil
}
