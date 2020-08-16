package event

import (
	"time"

	"github.com/google/uuid"
)

type Type string

const (
	EventTypeAction      Type = "action"
	EventTypeAlias       Type = "alias"
	EventTypeGroup       Type = "group"
	EventTypeIdentify    Type = "indentify"
	EventTypePageview    Type = "pageview"
	EventTypeScreenview  Type = "screenview"
	EventTypeSession     Type = "session"
	EventTypeTiming      Type = "timing"
	EventTypeTransaction Type = "transaction"
)

type Event struct {
	ID         string `json:"id" structs:"id"`
	TrackingID string `json:"trackingId" structs:"trackingID"`
	UserID     string `json:"userId,omitempty" structs:"userID,omitempty"`
	GroupID    string `json:"groupId,omitempty" structs:"groupID,omitempty"`
	SessionID  string `json:"sessionId,omitempty" structs:"sessionID,omitempty"`
	DeviceID   string `json:"deviceId,omitempty" structs:"deviceID,omitempty"`

	Event          Type `json:"event" structs:"event"`
	NonInteractive bool
	Channel        string    `json:"channel,omitempty" structs:"channel,omitempty"`
	Platform       string    `json:"platform,omitempty" structs:"platform,omitempty"`
	Timestamp      time.Time `json:"timestamp" structs:"timestamp"`
	Context        Contexts  `json:"context,omitempty" structs:"context,omitempty"`
}

func New(typ Type, opts ...Option) *Event {
	e := &Event{
		ID:        uuid.New().String(),
		Event:     typ,
		Timestamp: time.Now(),
		Context:   make(map[string]Context),
	}

	for _, o := range opts {
		o(e)
	}

	return e
}

func Empty() *Event {
	return &Event{
		ID:      uuid.New().String(),
		Context: make(map[string]Context),
	}
}

func (e *Event) SetContext(ctx Context) {
	e.Context[string(ctx.Type())] = ctx
}
