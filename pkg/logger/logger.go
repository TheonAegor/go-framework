package logger

import (
	"context"
)

var (
	DefaultLogger = NewLogger(nil)
	// DefaultLevel used by logger
	DefaultLevel Level = InfoLevel
)

type Loggerer interface {
	Debug(ctx context.Context, msg string, v ...interface{})
	Info(ctx context.Context, msg string, v ...interface{})
	Warning(ctx context.Context, msg string, v ...interface{})
	Error(ctx context.Context, msg string, v ...interface{})
	Fatal(ctx context.Context, msg string, v ...interface{})
}

type LoggerConfig struct {
	LogLevel string `json:"loglevel"`
}

func NewLogger(cfg *LoggerConfig) Loggerer {
	level := InfoLevel

	if cfg != nil {
		if lvl, err := GetLevel(cfg.LogLevel); err == nil {
			level = lvl
		}
	}
	opts := []Option{
		WithLevel(level),
	}

	return NewSlogLogger(opts...)
}
