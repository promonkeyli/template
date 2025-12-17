package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func Log(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		attrs := []slog.Attr{
			slog.String("method", c.Request.Method),
			slog.String("path", c.FullPath()),
			slog.Int("status", status),
			slog.Duration("latency", latency),
			slog.String("client_ip", c.ClientIP()),
			slog.String("user_agent", c.Request.UserAgent()),
		}

		if len(c.Errors) > 0 {
			attrs = append(attrs,
				slog.String("errors", c.Errors.String()),
			)
		}

		// slog.Attr → []any（关键修复点）
		args := make([]any, 0, len(attrs))
		for _, a := range attrs {
			args = append(args, a)
		}

		switch {
		case status >= 500:
			log.Error("http request error", args...)
		case status >= 400:
			log.Warn("http request warning", args...)
		default:
			log.Info("http request", args...)
		}
	}
}
