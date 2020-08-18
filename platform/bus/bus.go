package bus

import (
	"time"

	"github.com/blushft/strana"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/config"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/pkg/errors"
)

type Bus struct {
	conf config.Bus
	nats *natsBus

	pub    strana.Publisher
	sub    strana.Subscriber
	routes []*route

	started chan bool
	exit    chan struct{}
}

func New(conf config.Bus) (*Bus, error) {
	nats, err := newNatsBus(server.Options{
		Port:          conf.Port,
		HTTPPort:      conf.HTTPPort,
		Authorization: conf.Token,
	})

	if err != nil {
		return nil, err
	}

	return &Bus{
		conf:    conf,
		nats:    nats,
		started: make(chan bool),
		exit:    make(chan struct{}),
	}, nil
}

func (b *Bus) Publisher() strana.Publisher {
	return b.pub
}

func (b *Bus) Subscribe(topic string, fn func(*message.Message) error) error {
	sub, err := b.nats.NewSubscriber()
	if err != nil {
		return err
	}

	return sub.Subscribe(topic, fn)
}

func (b *Bus) Subscriber() strana.Subscriber {
	return b.sub
}

func (b *Bus) Mount(mod strana.Module) error {
	return mod.Events(b)
}

func (b *Bus) Handle(src, sink string, hndlr strana.EventHandlerFunc) error {
	conn, err := b.nats.newConn()
	if err != nil {
		return err
	}

	r, err := newRoute(src, sink, conn, hndlr)
	if err != nil {
		return err
	}

	b.routes = append(b.routes, r)

	return nil
}

func (b *Bus) Start() error {
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

func (b *Bus) Started() <-chan bool {
	return b.started
}

func (b *Bus) Shutdown() error {
	close(b.exit)
	b.nats.Shutdown()
	return nil
}
