package logger

import (
	"context"
)

type loggerKey struct{}

// FromContext returns logger from passed context
func FromContext(ctx context.Context) (Loggerer, bool) {
	if ctx == nil {
		return nil, false
	}
	l, ok := ctx.Value(loggerKey{}).(Loggerer)
	return l, ok
}

// NewContext stores logger into passed context
func NewContext(ctx context.Context, l Loggerer) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, loggerKey{}, l)
}

func FromOutgoingContext(ctx context.Context) Loggerer {
	if l, ok := ctx.Value(loggerKey{}).(Loggerer); ok {
		return l
	}
	return DefaultLogger
}

func FromIncomingContext(ctx context.Context) Loggerer {
	if l, ok := ctx.Value(loggerKey{}).(Loggerer); ok {
		return l
	}
	return DefaultLogger
}
