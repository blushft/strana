package log

import (
	"github.com/blushft/strana"
	"github.com/blushft/strana/domain/entity"
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

func (l *logger) Process(msg *entity.RawMessage) ([]*entity.RawMessage, error) {
	spew.Dump(msg)
	return []*entity.RawMessage{msg}, nil
}
