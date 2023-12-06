package logger

import (
	"context"
	"log/slog"
	"os"
)

type slogLog struct {
	l    *slog.Logger
	opts Options
}

func (s slogLog) Debug(ctx context.Context, msg string, v ...interface{}) {
	s.l.DebugContext(ctx, msg, v...)
}

func (s slogLog) Info(ctx context.Context, msg string, v ...interface{}) {
	s.l.InfoContext(ctx, msg, v...)
}

func (s slogLog) Warning(ctx context.Context, msg string, v ...interface{}) {
	s.l.WarnContext(ctx, msg, v...)
}

func (s slogLog) Error(ctx context.Context, msg string, v ...interface{}) {
	s.l.ErrorContext(ctx, msg, v...)
}

func (s slogLog) Fatal(ctx context.Context, msg string, v ...interface{}) {
	s.l.Log(ctx, loggerToSlogLevel(FatalLevel), msg, v...)
	os.Exit(1)
}

func NewSlogLogger(opts ...Option) Loggerer {
	sl := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	return slogLog{l: sl, opts: NewOptions(opts...)}
}

func loggerToSlogLevel(level Level) slog.Level {
	switch level {
	case DebugLevel:
		return slog.LevelDebug
	case WarnLevel:
		return slog.LevelWarn
	case ErrorLevel:
		return slog.LevelError
	case FatalLevel:
		return slog.LevelError + 1
	default:
		return slog.LevelInfo
	}
}
