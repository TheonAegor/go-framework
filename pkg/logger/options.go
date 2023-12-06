package logger

type Options struct {
	// The logging level the logger should log
	Level Level
}

// Option func
type Option func(*Options)

// NewOptions creates new options struct
func NewOptions(opts ...Option) Options {
	options := Options{
		Level: DefaultLevel,
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
