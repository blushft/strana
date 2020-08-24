package enhancer

import (
	"github.com/blushft/strana"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
	"github.com/blushft/strana/platform/store"
	"github.com/blushft/strana/processors"
	"github.com/gofiber/fiber"
	"github.com/mitchellh/mapstructure"
)

type Enhancer interface {
	strana.Module
}

type Options struct {
	Processors []config.Processor `json:"processors" yaml:"processors" mapstructure:"processors"`
}

type enhancer struct {
	conf config.Module
	opts Options
	log  *logger.Logger

	procs []strana.Processor
}

func New(conf config.Module) (Enhancer, error) {
	opts := Options{}
	if err := mapstructure.Decode(conf.Options, &opts); err != nil {
		return nil, err
	}

	procs := make([]strana.Processor, 0, len(opts.Processors))

	for _, p := range opts.Processors {
		proc, err := processors.New(p)
		if err != nil {
			return nil, err
		}

		procs = append(procs, proc)
	}

	return &enhancer{
		conf:  conf,
		opts:  opts,
		procs: procs,
	}, nil
}

func (e *enhancer) Routes(fiber.Router) error {
	return nil
}

func (e *enhancer) Services(*store.SQLStore) error {
	return nil
}

func (e *enhancer) Events(eh strana.EventHandler) error {
	return eh.Handle(e.conf.Source.Topic, e.conf.Sink.Topic, e.handle)
}

func (e *enhancer) Logger(l *logger.Logger) {
	e.log = l.WithFields(logger.Fields{"module": "enhancer"})
}

func (e *enhancer) handle(msg *message.Message) ([]*message.Message, error) {
	evts, err := processors.Execute(e.procs, msg.Event)
	if err != nil {
		return nil, err
	}

	results := make([]*message.Message, 0, len(evts))

	for _, en := range evts {
		results = append(results, message.NewMessage(en))
	}

	return results, nil
}
