package bus

import (
	"errors"

	"github.com/blushft/strana"
	"github.com/blushft/strana/platform"
	"github.com/blushft/strana/platform/bus/broker"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
	"github.com/oklog/run"
)

type Bus interface {
	strana.EventHandler
	Mount(strana.Module) error
	Start() error
	Shutdown() error
}

type bus struct {
	brokers map[string]broker.Broker
	routes  []Route

	stop chan struct{}

	log *logger.Logger
}

func New(conf config.Bus) (Bus, error) {
	brs := map[string]broker.Broker{}
	for _, c := range conf.Brokers {
		br, err := platform.NewBroker(c)
		if err != nil {
			return nil, err
		}

		brs[c.Name] = br
	}

	return &bus{
		brokers: brs,
		log:     logger.WithFields(logger.Fields{"component": "bus"}),
	}, nil
}

func (b *bus) Mount(mod strana.Module) error {
	return mod.Events(b)
}

func (b *bus) Handle(src, sink message.Path, h strana.EventHandlerFunc) error {
	panic("not implemented")
}

func (b *bus) Publisher() strana.Publisher {
	return b
}

func (b *bus) Publish(p message.Path, msg *message.Message) error {
	br, ok := b.brokers[p.Broker]
	if !ok {
		return errors.New("invalid broker")
	}

	return br.Publisher().Publish(p, msg)
}

func (b *bus) Subscribe(p message.Path, h strana.SubscriptionHandlerFunc) error {
	br, ok := b.brokers[p.Broker]
	if !ok {
		return errors.New("invalid broker")
	}

	return br.Subscriber().Subscribe(p, h)
}

func (b *bus) Subscriber() strana.Subscriber {
	return b
}

func (b *bus) Close() error {
	panic("not implemented")
}

func (b *bus) Start() error {
	g := run.Group{}

	for _, br := range b.brokers {
		g.Add(br.Connect, func(e error) {
			if err := br.Disconnect(); err != nil {
				b.log.WithError(err).Error("broker disconnect")
			}
		})
	}

	g.Add(func() error {
		<-b.stop
		return nil
	}, func(error) {
		b.log.Info("stopping bus")
	})

	return g.Run()
}

func (b *bus) Shutdown() error {
	close(b.stop)
	return nil
}
