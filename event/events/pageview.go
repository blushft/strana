package events

import (
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/event/contexts"
)

const EventTypePageview event.Type = "pageview"

func Pageview(hn string, opts ...event.Option) *event.Event {
	eopts := mergeOptions(
		[]event.Option{
			event.WithContext(&contexts.Page{Hostname: hn}),
			event.WithValidator(pageviewValidator()),
		},
		opts,
	)

	return event.New(EventTypePageview, eopts...)
}

func pageviewValidator() event.Validator {
	return event.NewValidator(
		event.WithRule("valid_page", event.ContextValid(contexts.ContextPage)),
	)
}
