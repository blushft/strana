package nats

import (
	"time"

	"github.com/blushft/strana"
	"github.com/blushft/strana/platform/bus"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/config"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/pkg/errors"
)

type natsbus struct {
	conf config.Bus
	nats *natsBus

	pub    strana.Publisher
	sub    strana.Subscriber
	routes []*bus.Route

	started chan bool
	exit    chan struct{}
}

func New(conf config.Bus) (*natsbus, error) {
	opts, err := unmarshalOptions(conf.Options)
	if err != nil {
		return nil, err
	}

	nats, err := newNatsBus(server.Options{
		Port:          opts.Port,
		HTTPPort:      opts.HTTPPort,
		Authorization: opts.Token,
	})

	if err != nil {
		return nil, err
	}

	return &natsbus{
		conf:    conf,
		nats:    nats,
		started: make(chan bool),
		exit:    make(chan struct{}),
	}, nil
}

func (b *natsbus) Publisher() strana.Publisher {
	return b.pub
}

func (b *natsbus) Subscribe(topic string, fn func(*message.Message) error) error {
	sub, err := b.nats.NewSubscriber()
	if err != nil {
		return err
	}

	return sub.Subscribe(topic, fn)
}

func (b *natsbus) Subscriber() strana.Subscriber {
	return b.sub
}

func (b *natsbus) Mount(mod strana.Module) error {
	return mod.Events(b)
}

func (b *natsbus) Handle(src, sink message.Path, hndlr strana.EventHandlerFunc) error {
	r, err := bus.NewRoute(src, sink, b, hndlr)
	if err != nil {
		return err
	}

	b.routes = append(b.routes, r)

	return nil
}

func (b *natsbus) Start() error {
	go b.nats.Start()

	if !b.nats.svr.ReadyForConnections(time.Second * 30) {
		return errors.New("unable to connect to nats bus")
	}

	pub, err := b.nats.NewPublisher()
	if err != nil {
		b.nats.Shutdown()
		return err
	}

	b.pub = pub
	defer b.pub.Close()

	sub, err := b.nats.NewSubscriber()
	if err != nil {
		b.nats.Shutdown()

		return err
	}

	b.sub = sub
	defer b.sub.Close()

	b.started <- true
	<-b.exit

	return nil
}

func (b *natsbus) Started() <-chan bool {
	return b.started
}

func (b *natsbus) Shutdown() error {
	close(b.exit)
	b.nats.Shutdown()

	return nil
}
