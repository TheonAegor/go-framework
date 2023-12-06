package codec

// Codec encodes/decodes various types of messages.
type Codec interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
}

// Message represents detailed information about
// the communication, likely followed by the body.
// In the case of an error, body may be nil.
type Message struct {
	Id       string
	Type     MessageType
	Target   string
	Method   string
	Endpoint string
	Error    string

	Body []byte
}

// MessageType specifies message type for codec
type MessageType int
