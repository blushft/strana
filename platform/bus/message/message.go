package message

import (
	"github.com/blushft/strana/event"
	"github.com/google/uuid"
)

type UUIDGenerator func() uuid.UUID

var uuidGenerator = func() uuid.UUID {
	return uuid.New()
}

func SetUUIDGenerator(fn UUIDGenerator) {
	uuidGenerator = fn
}

type Message struct {
	ID       uuid.UUID
	Metadata map[string]string
	Event    *event.Event

	state string
}

type MessageOption func(*Message)

func WithMetadata(k, v string) MessageOption {
	return func(m *Message) {
		m.Metadata[k] = v
	}
}

func SetMetadata(meta map[string]string) MessageOption {
	return func(m *Message) {
		m.Metadata = meta
	}
}

func NewMessage(evt *event.Event, opts ...MessageOption) *Message {
	msg := &Message{
		ID:       uuidGenerator(),
		Metadata: make(map[string]string),
		Event:    evt,
	}

	for _, o := range opts {
		o(msg)
	}

	return msg
}

func (msg *Message) Envelope() (Envelope, error) {
	return NewEnvelope(msg)
}
