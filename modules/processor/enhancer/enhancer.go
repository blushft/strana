package enhancer

import (
	"github.com/blushft/strana"
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/modules"
	"github.com/blushft/strana/platform"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
	"github.com/blushft/strana/platform/store"
	"github.com/blushft/strana/processor"
	"github.com/gofiber/fiber"
	"github.com/mitchellh/mapstructure"
)

func init() {
	modules.Register("enhancer", New)
}

type Enhancer interface {
	strana.Processor
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

	procs []processor.EventProcessor
}

func New(conf config.Module) (strana.Module, error) {
	opts := Options{}
	if err := mapstructure.Decode(conf.Options, &opts); err != nil {
		return nil, err
	}

	procs := make([]processor.EventProcessor, 0, len(opts.Processors))

	for _, p := range opts.Processors {
		proc, err := platform.NewEventProcessor(p)
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

func (mod *enhancer) Routes(fiber.Router) error {
	return nil
}

func (mod *enhancer) Services(*store.SQLStore) error {
	return nil
}

func (mod *enhancer) Events(eh strana.EventHandler) error {
	mod.pub = eh.Publisher()
	mod.sub = eh.Subscriber()

	return eh.Handle(mod.conf.Source, mod.conf.Sink, mod.handle)
}

func (mod *enhancer) Logger(l *logger.Logger) {
	mod.log = l.WithFields(logger.Fields{"module": "enhancer"})
}

func (mod *enhancer) Publish(evt *event.Event) error {
	return mod.pub.Publish(mod.conf.Source, message.NewMessage(evt))
}

func (mod *enhancer) Subscribe(fn strana.SubscriptionHandlerFunc) error {
	return mod.sub.Subscribe(mod.conf.Source, fn)
}

func (mod *enhancer) handle(msg *message.Message) ([]*message.Message, error) {
	evts, err := processor.Execute(mod.procs, msg.Event)
	if err != nil {
		return nil, err
	}

	results := make([]*message.Message, 0, len(evts))

	for _, en := range evts {
		results = append(results, message.NewMessage(en))
	}

	return results, nil
}
