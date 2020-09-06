package events

import (
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/event/contexts"
)

const EventTypeSession event.Type = "session"

func Session(id string, opts ...event.Option) *event.Event {
	eopts := mergeOptions(
		[]event.Option{
			event.WithContext(&contexts.Session{ID: id}),
			event.WithValidator(sessionValidator()),
		},
		opts,
	)

	return event.New(EventTypeSession, eopts...)
}

func sessionValidator() event.Validator {
	return event.NewValidator(
		event.WithRule("valid_session", event.ContextValid(contexts.ContextSession)),
	)
}
