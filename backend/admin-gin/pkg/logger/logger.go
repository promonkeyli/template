package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"time"
)

// 定义 Context 中 TraceID 的 Key
type ctxKey string

const (
	TraceIDKey ctxKey = "trace_id"
)

type Options struct {
	Level  string    // 日志级别: debug, info, warn, error
	Format string    // 输出格式: json, text
	Output io.Writer // 输出位置: 默认为 os.Stdout，也可以是文件
}

// Init 初始化全局 Logger
func Init(opts Options) {
	// 1. 设置日志级别
	var level slog.Level
	switch opts.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	// 2. 自定义 HandlerOptions
	handlerOpts := &slog.HandlerOptions{
		Level:     level,
		AddSource: true, // 生产环境如果对性能要求极高，可以关闭
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// 最佳实践：统一时间格式
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					return slog.String(slog.TimeKey, t.Format(time.DateTime))
				}
			}
			return a
		},
	}

	// 3. 选择输出格式
	var handler slog.Handler
	output := opts.Output
	if output == nil {
		output = os.Stdout
	}

	if opts.Format == "json" {
		handler = slog.NewJSONHandler(output, handlerOpts)
	} else {
		handler = slog.NewTextHandler(output, handlerOpts)
	}

	// 4. 包装 ContextHandler (这是关键，用于自动提取 TraceID)
	ctxHandler := &ContextHandler{handler}

	// 5. 设置为全局默认 Logger
	logger := slog.New(ctxHandler)
	slog.SetDefault(logger)
}

// ---------------------------------------------------------
// 自定义 Handler 用于从 Context 提取 TraceID
// ---------------------------------------------------------

type ContextHandler struct {
	slog.Handler
}

// Handle 重写 Handle 方法，从 context 中获取 trace_id 并注入日志
func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if traceID, ok := ctx.Value(TraceIDKey).(string); ok && traceID != "" {
		r.AddAttrs(slog.String("trace_id", traceID))
	}
	return h.Handler.Handle(ctx, r)
}

// WithAttrs 必须重写，确保返回的是 ContextHandler
func (h *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ContextHandler{h.Handler.WithAttrs(attrs)}
}

// WithGroup 必须重写
func (h *ContextHandler) WithGroup(name string) slog.Handler {
	return &ContextHandler{h.Handler.WithGroup(name)}
}

// ---------------------------------------------------------
// 辅助函数：由于 slog.Default() 已经是全局的，我们直接封装几个快捷方式
// 但推荐直接使用 slog.InfoContext 等原生方法
// ---------------------------------------------------------

// WithTraceID 将 traceID 注入到 context 中 (用于中间件)
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceID)
}

// LogError 是一个辅助函数，用于标准化的记录 error
// 最佳实践：总是同时记录 error 消息和堆栈（如果 error 包支持）或上下文
func LogError(ctx context.Context, msg string, err error, attrs ...slog.Attr) {
	newAttrs := append([]slog.Attr{slog.String("error", err.Error())}, attrs...)
	slog.ErrorContext(ctx, msg, anyList(newAttrs)...)
}

func anyList(attrs []slog.Attr) []any {
	args := make([]any, len(attrs))
	for i, a := range attrs {
		args[i] = a
	}
	return args
}
