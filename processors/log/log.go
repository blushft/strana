package log

import (
	"github.com/blushft/strana"
	"github.com/blushft/strana/pkg/event"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/processors"
	"github.com/davecgh/go-spew/spew"
)

func init() {
	processors.Register("log", func(config.Processor) (strana.Processor, error) {
		return &logger{}, nil
	})
}

type logger struct{}

func (l *logger) Process(evt *event.Event) ([]*event.Event, error) {
	spew.Dump(evt)
	return []*event.Event{evt}, nil
}
