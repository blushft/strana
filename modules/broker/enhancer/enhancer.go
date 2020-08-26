package enhancer

import (
	"github.com/blushft/strana"
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/modules"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
	"github.com/blushft/strana/platform/store"
	"github.com/blushft/strana/processors"
	"github.com/gofiber/fiber"
	"github.com/mitchellh/mapstructure"
)

func init() {
	modules.Register("enhancer", New)
}

type Enhancer interface {
	strana.Broker
}

type Options struct {
	Processors []config.Processor `json:"processors" yaml:"processors" mapstructure:"processors"`
}

type enhancer struct {
	conf config.Module
	opts Options
	log  *logger.Logger

	pub strana.Publisher
	sub strana.Subscriber

	procs []strana.Processor
}

func New(conf config.Module) (strana.Module, error) {
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
	e.pub = eh.Publisher()
	e.sub = eh.Subscriber()

	return eh.Handle(e.conf.Source.Topic, e.conf.Sink.Topic, e.handle)
}

func (e *enhancer) Logger(l *logger.Logger) {
	e.log = l.WithFields(logger.Fields{"module": "enhancer"})
}

func (e *enhancer) Publish(evt *event.Event) error {
	return e.pub.Publish(e.conf.Source.Topic, message.NewMessage(evt))
}

func (e *enhancer) Subscribe(fn strana.SubscriptionHandlerFunc) error {
	return e.sub.Subscribe(e.conf.Source.Topic, fn)
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
