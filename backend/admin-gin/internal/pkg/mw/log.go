package mw

import (
	"log/slog"
	"mall-api/internal/logger"
	"time"

	"github.com/gin-gonic/gin"
)

// SlogLog 是简单的 slog 中间件
func SlogLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		logger.Log.Info("request",
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.Int("status", c.Writer.Status()),
			slog.String("ip", c.ClientIP()),
			slog.Duration("duration", time.Since(start)),
		)
	}
}
