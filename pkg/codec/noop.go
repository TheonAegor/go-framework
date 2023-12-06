package codec

import (
	"encoding/json"
)

type noopCodec struct{}

// Frame gives us the ability to define raw data to send over the pipes
type Frame struct {
	Data []byte
}

func (c *noopCodec) Marshal(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, nil
	}

	switch ve := v.(type) {
	case string:
		return []byte(ve), nil
	case *string:
		return []byte(*ve), nil
	case []byte:
		return ve, nil
	case *[]byte:
		return *ve, nil
	case *Frame:
		return ve.Data, nil
	case *Message:
		return ve.Body, nil
	}

	return json.Marshal(v)
}

func (c *noopCodec) Unmarshal(d []byte, v interface{}) error {
	if v == nil {
		return nil
	}
	switch ve := v.(type) {
	case *string:
		*ve = string(d)
		return nil
	case []byte:
		copy(ve, d)
		return nil
	case *[]byte:
		*ve = d
		return nil
	case *Frame:
		ve.Data = d
		return nil
	case *Message:
		ve.Body = d
		return nil
	}

	return json.Unmarshal(d, v)
}

// NewCodec returns new noop codec
func NewCodec() Codec {
	return &noopCodec{}
}
