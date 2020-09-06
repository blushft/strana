package events

import (
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/event/contexts"
)

const EventTypeTiming event.Type = "timing"

func Timing(t *contexts.Timing, opts ...event.Option) *event.Event {
	eopts := mergeOptions(
		[]event.Option{
			event.WithContext(t),
			event.WithValidator(timingValidator()),
		},
		opts,
	)

	return event.New(EventTypeTiming, eopts...)
}

func timingValidator() event.Validator {
	return event.NewValidator(
		event.WithRule("valid_timing", event.ContextValid(contexts.ContextTiming)),
	)
}
