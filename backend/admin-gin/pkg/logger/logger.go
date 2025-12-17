package logger

import (
	"log/slog"
	"os"
)

var baseLogger *slog.Logger

// Init 初始化全局 logger（只调用一次）
func Init(cfg Config) {
	opts := &slog.HandlerOptions{
		Level: cfg.Level,
	}

	var handler slog.Handler
	if cfg.Env == "dev" {
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	baseLogger = slog.New(handler).With(
		"service", cfg.Service,
		"env", cfg.Env,
	)

	slog.SetDefault(baseLogger)
}

// Default 返回全局 logger
func Default() *slog.Logger {
	if baseLogger == nil {
		return slog.Default()
	}
	return baseLogger
}
