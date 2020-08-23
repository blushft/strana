package log

import (
	"encoding/json"

	"github.com/blushft/strana"
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
	"github.com/blushft/strana/processors"
)

func init() {
	processors.Register("log", func(config.Processor) (strana.Processor, error) {
		return &loggerp{}, nil
	})
}

type loggerp struct{}

func (l *loggerp) Process(evt *event.Event) ([]*event.Event, error) {
	b, err := json.MarshalIndent(evt, "  ", "  ")
	if err != nil {
		logger.Log().WithError(err).Error("unable to marshal event")
	} else {
		logger.Log().Info(string(b))
	}
	return []*event.Event{evt}, nil
}
