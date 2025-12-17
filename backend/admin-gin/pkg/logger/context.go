package logger

import (
	"context"
	"log/slog"
)

type ctxKey struct{}

// WithContext 将 logger 写入 context
func WithContext(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, l)
}

// FromContext 从 context 获取 logger（推荐使用）
func FromContext(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*slog.Logger); ok {
		return l
	}
	return Default()
}
