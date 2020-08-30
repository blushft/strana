package fanin

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
)

func init() {
	modules.Register("fanin", New)
}

type FanIn interface {
	strana.Broker
}

type Options struct {
	Sources    []config.MessagePath
	Processors []config.Processor
}

type fanIn struct {
	conf config.Module
	opts Options
	log  *logger.Logger

	pub strana.Publisher
	sub strana.Subscriber

	srcs  []config.MessagePath
	procs []strana.Processor
}

func New(conf config.Module) (strana.Module, error) {
	opts := Options{}
	if err := modules.BindOptions(conf.Options, &opts); err != nil {
		return nil, err
	}

	procs, err := processors.NewSet(opts.Processors)
	if err != nil {
		return nil, err
	}

	return &fanIn{
		conf:  conf,
		opts:  opts,
		procs: procs,
	}, nil
}

func (mod *fanIn) Routes(fiber.Router) error {
	return nil
}

func (mod *fanIn) Services(*store.SQLStore) error {
	return nil
}

func (mod *fanIn) Events(eh strana.EventHandler) error {
	mod.pub = eh.Publisher()
	mod.sub = eh.Subscriber()

	for _, src := range mod.srcs {
		if err := eh.Handle(src.Topic, mod.conf.Source.Topic, mod.handle); err != nil {
			return err
		}
	}

	return nil
}

func (mod *fanIn) Logger(l *logger.Logger) {
	mod.log = l.WithFields(logger.Fields{"module": "fan_in_broker"})
}

func (mod *fanIn) Publish(evt *event.Event) error {
	return mod.pub.Publish(mod.conf.Source.Topic, message.NewMessage(evt))
}

func (mod *fanIn) Subscribe(fn strana.SubscriptionHandlerFunc) error {
	return mod.sub.Subscribe(mod.conf.Source.Topic, fn)
}

func (mod *fanIn) handle(msg *message.Message) ([]*message.Message, error) {
	evts, err := processors.Execute(mod.procs, msg.Event)
	if err != nil {
		return nil, err
	}

	results := make([]*message.Message, 0, len(evts))

	for _, en := range evts {
		results = append(results, message.NewMessage(en))
	}

	return results, nil
}
