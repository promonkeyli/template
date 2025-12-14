package mw

import (
	"log/slog"
	"net/http"
	"runtime/debug"

	"mall-api/internal/logger"

	"github.com/gin-gonic/gin"
)

func SlogRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Error("panic recovered",
					slog.Any("error", err),
					slog.String("stack", string(debug.Stack())),
				)

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Internal Server Error",
				})
			}
		}()

		c.Next()
	}
}
