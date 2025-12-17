package logger

import "log/slog"

type Config struct {
	Service string
	Env     string // dev / test / prod
	Level   slog.Level
}
