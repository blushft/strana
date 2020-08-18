package message

import (
	"encoding/json"
)

type Marshaler func(*Message) (Envelope, error)

var marshaler = func(msg *Message) (Envelope, error) {
	return json.Marshal(msg)
}

func SetMarshaler(fn Marshaler) {
	marshaler = fn
}

type Unmarshaler func([]byte) (*Message, error)

var unmarshaler = func(b []byte) (*Message, error) {
	var msg *Message
	if err := json.Unmarshal(b, &msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func SetUnmarshaler(fn Unmarshaler) {
	unmarshaler = fn
}

type Envelope []byte

func NewEnvelope(msg *Message) (Envelope, error) {
	return marshaler(msg)
}

func (e Envelope) Copy() Envelope {
	var ec Envelope
	copy(ec, e)

	return ec
}

func (e Envelope) Unmarshal() (*Message, error) {
	return unmarshaler(e)
}
