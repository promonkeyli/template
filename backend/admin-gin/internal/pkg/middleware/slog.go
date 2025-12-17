// 基于 slog 的 gin 日志中间件
package middleware

import (
	"mall-api/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Slog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		reqID := uuid.NewString()

		l := logger.Default().With(
			logger.FieldRequestID, reqID,
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
		)

		ctx := logger.WithContext(c.Request.Context(), l)
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		l.Info("http request completed",
			"status", c.Writer.Status(),
			"latency_ms", time.Since(start).Milliseconds(),
		)
	}
}
