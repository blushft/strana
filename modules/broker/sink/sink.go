package sink

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
	modules.Register("sink", New)
}

type Sink interface {
	strana.Broker
}

type Options struct {
	Destination config.Module
	Processors  []config.Processor
}

type sinkBroker struct {
	conf config.Module
	opts Options
	log  *logger.Logger

	pub strana.Publisher
	sub strana.Subscriber

	sink  strana.Sink
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

	return &sinkBroker{
		conf:  conf,
		opts:  opts,
		procs: procs,
	}, nil
}

func (mod *sinkBroker) Routes(fiber.Router) error {
	return nil
}

func (mod *sinkBroker) Services(*store.SQLStore) error {
	return nil
}

func (mod *sinkBroker) Events(eh strana.EventHandler) error {
	mod.pub = eh.Publisher()
	mod.sub = eh.Subscriber()

	return nil
}

func (mod *sinkBroker) Logger(l *logger.Logger) {
	mod.log = l.WithFields(logger.Fields{"module": "sink_broker"})
}

func (mod *sinkBroker) Publish(evt *event.Event) error {
	return mod.pub.Publish(mod.conf.Source.Topic, message.NewMessage(evt))
}

func (mod *sinkBroker) Subscribe(fn strana.SubscriptionHandlerFunc) error {
	return mod.sub.Subscribe(mod.conf.Source.Topic, fn)
}

func (mod *sinkBroker) handle(msg *message.Message) ([]*message.Message, error) {
	evts, err := processors.Execute(mod.procs, msg.Event)
	if err != nil {
		return nil, err
	}

	results := make([]*message.Message, 0, len(evts))

	for _, evt := range evts {
		if err := mod.sink.Publish(evt); err != nil {
			return nil, err
		}

		results = append(results, message.NewMessage(evt))
	}

	return results, nil
}
