package logger

import (
	"io"
	"log/slog"
	"os"
)

func NewLog(cfg Config) *slog.Logger {
	var writer io.Writer

	fileWriter := newWriter(cfg)

	// 开发环境：同时输出到 stdout
	if cfg.Format == "text" {
		writer = io.MultiWriter(fileWriter, os.Stdout)
	} else {
		writer = fileWriter
	}

	opts := &slog.HandlerOptions{
		Level: parseLevel(cfg.Level),
	}

	var handler slog.Handler
	switch cfg.Format {
	case "json":
		handler = slog.NewJSONHandler(writer, opts)
	default:
		handler = slog.NewTextHandler(writer, opts)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)

	return logger
}
