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
	ID         string `json:"id" structs:"id" mapstructure:"id"`
	TrackingID string `json:"trackingId" structs:"trackingID" mapstructure:"trackingID"`
	UserID     string `json:"userId,omitempty" structs:"userID,omitempty" mapstructure:"userID,omitempty"`
	Anonymous  bool   `json:"anonymous" structs:"anonymous" mapstructure:"anonymous"`
	GroupID    string `json:"groupId,omitempty" structs:"groupID,omitempty" mapstructure:"groupID,omitempty"`
	SessionID  string `json:"sessionId,omitempty" structs:"sessionID,omitempty" mapstructure:"sessionID,omitempty"`
	DeviceID   string `json:"deviceId,omitempty" structs:"deviceID,omitempty" mapstructure:"deviceID,omitempty"`

	Event          Type      `json:"event" structs:"event" mapstructure:"event"`
	NonInteractive bool      `json:"nonInteractive,omitempty" structs:"nonInteractive,omitempty" mapstructure:"nonInteractive,omitempty"`
	Channel        string    `json:"channel,omitempty" structs:"channel,omitempty" mapstructure:"channel,omitempty"`
	Platform       string    `json:"platform,omitempty" structs:"platform,omitempty" mapstructure:"platform,omitempty"`
	Timestamp      time.Time `json:"timestamp" structs:"timestamp" mapstructure:"timestamp"`
	Context        Contexts  `json:"context,omitempty" structs:"context,omitempty" mapstructure:"context,omitempty"`
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
