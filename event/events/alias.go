package events

import (
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/event/contexts"
)

const EventTypeAlias event.Type = "alias"

func Alias(from, to, user string, opts ...event.Option) *event.Event {
	eopts := mergeOptions(
		[]event.Option{
			event.WithContext(&contexts.User{UserID: user}),
			event.WithContext(&contexts.Alias{From: from, To: to}),
			event.WithValidator(aliasValidator()),
		},
		opts,
	)

	return event.New(EventTypeAlias, eopts...)
}

func aliasValidator() event.Validator {
	return event.NewValidator(
		event.WithRule("valid_group", event.ContextValid(contexts.ContextAlias)),
		event.WithRule("valid_user", event.ContextValid(contexts.ContextUser)),
	)
}
