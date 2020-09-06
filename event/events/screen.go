package events

import (
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/event/contexts"
)

const EventTypeScreen event.Type = "screen"

func Screen(name string, opts ...event.Option) *event.Event {
	eopts := mergeOptions(
		[]event.Option{
			event.WithContext(&contexts.Screen{Name: name}),
			event.WithValidator(screenValidator()),
		},
		opts,
	)

	return event.New(EventTypeScreen, eopts...)
}

func screenValidator() event.Validator {
	return event.NewValidator(
		event.WithRule("valid_screen", event.ContextValid(contexts.ContextScreen)),
	)
}
