package logger

import (
	"io"
	"log/slog"
	"os"
)

var Log *slog.Logger

// 用来适配 io.Writer → slog
type slogWriter struct{}

func (w *slogWriter) Write(p []byte) (n int, err error) {
	Log.Info(string(p))
	return len(p), nil
}
func Writer() io.Writer {
	return &slogWriter{}
}

func Init() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	Log = slog.New(handler)
}
