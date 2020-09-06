package events

import (
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/event/contexts"
)

const EventTypeGroup event.Type = "group"

func Group(name string, user string, opts ...event.Option) *event.Event {
	eopts := mergeOptions(
		[]event.Option{
			event.WithContext(&contexts.User{UserID: user}),
			event.WithContext(&contexts.Group{Name: name}),
			event.WithValidator(groupValidator()),
		},
		opts,
	)

	return event.New(EventTypeAction, eopts...)
}

func groupValidator() event.Validator {
	return event.NewValidator(
		event.WithRule("group_context", event.HasContext(contexts.ContextGroup)),
		event.WithRule("user_context", event.HasContext(contexts.ContextUser)),
	)
}
