package log

import (
	"log"

	"github.com/blushft/strana"
	"github.com/blushft/strana/domain/entity"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/processors"
)

func init() {
	processors.Register("log", func(config.Processor) (strana.Processor, error) {
		return &logger{}, nil
	})
}

type logger struct{}

func (l *logger) Process(msg *entity.RawMessage) ([]*entity.RawMessage, error) {
	log.Printf("%v", msg)

	return []*entity.RawMessage{msg}, nil
}
