package tracer

type noopTracer struct {
	opts Options
}

func NewTracer(opts ...Option) Tracer {
	return &noopTracer{
		opts: NewOptions(opts...),
	}
}
