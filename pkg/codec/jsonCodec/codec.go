// Package json provides a json codec
package jsonCodec

import (
	"encoding/json"
	"github.com/TheonAegor/go-framework/pkg/codec"
	"io"
	"io/ioutil"
)

type jsonCodec struct{}

func (c *jsonCodec) Marshal(b interface{}) ([]byte, error) {
	switch m := b.(type) {
	case nil:
		return nil, nil
	case *codec.Frame:
		return m.Data, nil
	}

	return json.Marshal(b)
}

func (c *jsonCodec) Unmarshal(b []byte, v interface{}) error {
	if b == nil {
		return nil
	}
	switch m := v.(type) {
	case nil:
		return nil
	case *codec.Frame:
		m.Data = b
		return nil
	}

	return json.Unmarshal(b, v)
}

func (c *jsonCodec) ReadHeader(conn io.Reader, m *codec.Message, t codec.MessageType) error {
	return nil
}

func (c *jsonCodec) ReadBody(conn io.Reader, b interface{}) error {
	switch m := b.(type) {
	case nil:
		return nil
	case *codec.Frame:
		buf, err := ioutil.ReadAll(conn)
		if err != nil {
			return err
		}
		m.Data = buf
		return nil
	}

	return json.NewDecoder(conn).Decode(b)
}

func (c *jsonCodec) Write(conn io.Writer, m *codec.Message, b interface{}) error {
	switch m := b.(type) {
	case nil:
		return nil
	case *codec.Frame:
		_, err := conn.Write(m.Data)
		return err
	}

	return json.NewEncoder(conn).Encode(b)
}

func (c *jsonCodec) String() string {
	return "json"
}

func NewCodec() codec.Codec {
	return &jsonCodec{}
}
