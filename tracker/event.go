package tracker

import (
	"encoding/json"
	"time"

	"github.com/fatih/structs"
	"github.com/google/uuid"
)

type EventType string

const (
	EventTypeAction      EventType = "action"
	EventTypePageview              = "pageview"
	EventTypeScreenview            = "screenview"
	EventTypeSession               = "session"
	EventTypeGroup                 = "group"
	EventTypeTransaction           = "transaction"
	EventTypeTimed                 = "timed_event"
)

type Event struct {
	ID         string `json:"eid" mapstructure:"id"`
	AppID      int    `json:"aid" mapstructure:"aid"`
	TrackingID string `json:"tid" mapstructure:"tid"`
	UserID     string `json:"uid" mapstructure:"uid"`
	GroupID    string `json:"gid" mapstructure:"gid"`
	DeviceID   string `json:"cid" mapstructure:"cid"`
	SessionID  string `json:"sid" mapstructure:"sid"`

	Type           string    `json:"e" mapstructure:"e"`
	NonInteractive bool      `json:"ni" mapstructure:"ni"`
	Datasource     string    `json:"ds" mapstructure:"ds"`
	Namespace      string    `json:"ns" mapstructure:"ns"`
	Platform       string    `json:"p" mapstructure:"p"`
	AppVersion     string    `json:"av" mapstructure:"av"`
	Timestamp      time.Time `json:"dtm" mapstructure:"dtm"`

	Contexts map[string]Context `json:"-" mapstructure:"-"`
}

type EventOption func(*Event)

func NewEvent(typ EventType, opts ...EventOption) *Event {
	evt := &Event{
		ID:        uuid.New().String(),
		Type:      string(typ),
		Contexts:  make(map[string]Context),
		Timestamp: time.Now(),
	}

	for _, o := range opts {
		o(evt)
	}

	return evt
}

func AppID(id int) EventOption {
	return func(e *Event) {
		e.AppID = id
	}
}

func TrackingID(id string) EventOption {
	return func(e *Event) {
		e.TrackingID = id
	}
}

func UserID(id string) EventOption {
	return func(e *Event) {
		e.UserID = id
	}
}

func GroupID(id string) EventOption {
	return func(e *Event) {
		e.GroupID = id
	}
}

func DeviceID(id string) EventOption {
	return func(e *Event) {
		e.DeviceID = id
	}
}

func SessionID(id string) EventOption {
	return func(e *Event) {
		e.SessionID = id
	}
}

func NonInteractive() EventOption {
	return func(e *Event) {
		e.NonInteractive = true
	}
}

func Interactive() EventOption {
	return func(e *Event) {
		e.NonInteractive = false
	}
}

func Datasource(s string) EventOption {
	return func(e *Event) {
		e.Datasource = s
	}
}

func Namespace(s string) EventOption {
	return func(e *Event) {
		e.Namespace = s
	}
}

func Platform(s string) EventOption {
	return func(e *Event) {
		e.Platform = s
	}
}

func AppVersion(s string) EventOption {
	return func(e *Event) {
		e.AppVersion = s
	}
}

func WithContext(n string, c Context) EventOption {
	return func(e *Event) {
		e.Contexts[n] = c
	}
}

func WithDeviceContext(d *Device) EventOption {
	return WithContext("device", d)
}

func WithActionContext(a *Action) EventOption {
	return WithContext("action", a)
}

func (e *Event) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{}

	em := structs.Map(e)
	for k, v := range em {
		m[k] = v
	}

	for _, ctx := range e.Contexts {
		tctx := ctx
		for k, v := range tctx.Values() {
			m[k] = v
		}
	}

	return json.Marshal(m)
}
