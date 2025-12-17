package logger

import (
	"io"
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

func newWriter(cfg Config) io.Writer {
	if cfg.Dir != "" {
		_ = os.MkdirAll(cfg.Dir, 0755)
	}

	return &lumberjack.Logger{
		Filename: filepath.Join(cfg.Dir, cfg.Filename),
		MaxSize:  cfg.MaxSize, // MB
		MaxAge:   cfg.MaxAge,  // days
		Compress: cfg.Compress,
	}
}
