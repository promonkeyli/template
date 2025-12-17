package logger

import (
	"io"
	"log/slog"
	"os"
)

func NewLog(cfg Config) *slog.Logger {
	var writer io.Writer

	fileWriter := newWriter(cfg)

	writer = io.MultiWriter(fileWriter, os.Stdout)

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
