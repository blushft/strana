package events

import (
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/event/contexts"
)

const EventTypeAction event.Type = "action"

func Action(cat, action string, opts ...event.Option) *event.Event {
	eopts := mergeOptions(
		[]event.Option{
			event.WithContext(contexts.NewAction(cat, action)),
			event.WithValidator(actionValidator()),
		},
		opts,
	)

	return event.New(EventTypeAction, eopts...)
}

func actionValidator() event.Validator {
	return event.NewValidator(
		event.WithRule("valid_action", event.ContextValid(contexts.ContextAction)),
	)
}
