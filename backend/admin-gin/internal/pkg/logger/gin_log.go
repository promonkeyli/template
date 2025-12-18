// gin 日志适配器
package logger

import (
	"context"
	"io"
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
)

func BuilderGinLog(logger *slog.Logger) {

	// 1. 禁用 Gin 的控制台彩色输出（避免 JSON 中出现乱码字符）
	gin.DisableConsoleColor()

	// 2. 定义一个内部使用的 Writer 转换器
	// 这样逻辑就集中在这个函数里，不会暴露到外面
	newSlogWriter := func(level slog.Level) io.Writer {
		return &slogWriterAdapter{
			logger: logger,
			level:  level,
		}
	}

	// 3. 彻底接管 Gin 的全局 Writer
	gin.DefaultWriter = newSlogWriter(slog.LevelInfo)
	gin.DefaultErrorWriter = newSlogWriter(slog.LevelError)
}

// 私有适配器结构体
type slogWriterAdapter struct {
	logger *slog.Logger
	level  slog.Level
}

func (w *slogWriterAdapter) Write(p []byte) (n int, err error) {
	msg := strings.TrimSpace(string(p))
	if msg != "" {
		w.logger.Log(context.Background(), w.level, msg)
	}
	return len(p), nil
}
