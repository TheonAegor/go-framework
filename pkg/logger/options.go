package logger

import "os"

type Options struct {
	// The logging level the logger should log
	Level Level
	Out   *os.File
}

// Option func
type Option func(*Options)

// NewOptions creates new options struct
func NewOptions(opts ...Option) Options {
	options := Options{
		Level: DefaultLevel,
		Out:   os.Stdout,
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

// WithLevel set default level for the logger
func WithLevel(level Level) Option {
	return func(o *Options) {
		o.Level = level
	}
}

func WithFile(f *os.File) Option {
	return func(o *Options) {
		o.Out = f
	}
}
