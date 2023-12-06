package tracer

var (
	DefaultTracer Tracer = NewTracer()
)

type Tracer interface {
}
