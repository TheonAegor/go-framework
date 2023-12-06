package meter

// NoopMeter is an noop implementation of Meter
type noopMeter struct {
}

// NewMeter returns a configured noop reporter:
func NewMeter() Meter {
	return &noopMeter{}
}
