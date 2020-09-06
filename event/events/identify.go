package events

import (
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/event/contexts"
)

const EventTypeIdentify event.Type = "identify"

func Identify(user string, opts ...event.Option) *event.Event {
	eopts := mergeOptions(
		[]event.Option{
			event.WithContext(&contexts.User{UserID: user}),
			event.WithValidator(identifyValidator()),
		},
		opts,
	)

	return event.New(EventTypeIdentify, eopts...)
}

func identifyValidator() event.Validator {
	return event.NewValidator(
		event.WithRule("valid_user", event.ContextValid(contexts.ContextUser)),
	)
}
