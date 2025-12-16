package middleware

import (
	"log/slog"
	"mall-api/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// 1. 生成或获取 TraceID
		traceID := c.GetHeader("X-Request-Id")
		if traceID == "" {
			traceID = uuid.New().String()
		}

		// 2. 将 TraceID 注入 Request 的 Context 中
		// 注意：c.Request.Context() 和 c 是两回事，slog 用的是 context.Context
		ctx := logger.WithTraceID(c.Request.Context(), traceID)
		c.Request = c.Request.WithContext(ctx)

		// 同时也设置 Response Header
		c.Header("X-Request-Id", traceID)

		// 处理请求
		c.Next()

		// 3. 请求处理完后记录访问日志
		cost := time.Since(start)

		// 组装一些通用的 HTTP 属性
		status := c.Writer.Status()
		method := c.Request.Method
		clientIP := c.ClientIP()

		// 使用 slog.LogAttrs 可以避免反射，性能更高
		// 注意：这里一定要传入 c.Request.Context()，因为那里包含 trace_id
		level := slog.LevelInfo
		if status >= 500 {
			level = slog.LevelError
		}

		slog.LogAttrs(c.Request.Context(), level, "http_request",
			slog.Int("status", status),
			slog.String("method", method),
			slog.String("path", path),
			slog.String("query", raw),
			slog.String("ip", clientIP),
			slog.String("user_agent", c.Request.UserAgent()),
			slog.Duration("cost", cost),
		)
	}
}
