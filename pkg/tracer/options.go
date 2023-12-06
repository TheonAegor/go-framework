package tracer

// Options struct
type Options struct {
}

// Option func signature
type Option func(o *Options)

// NewOptions returns default options
func NewOptions(opts ...Option) Options {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}
	return options
}
