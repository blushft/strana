package nsq

import (
	"github.com/blushft/strana"
	"github.com/blushft/strana/platform/bus"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/nsqio/go-nsq"
	"github.com/nsqio/nsq/nsqd"
)

type nsqEmbedded struct {
	opts *nsqd.Options
	nsqd *nsqd.NSQD

	stop chan struct{}
}

func NewDefault(opts ...bus.Option) bus.Bus {
	b, err := New(opts...)
	if err != nil {
		panic(err)
	}

	return b
}

func New(opts ...bus.Option) (bus.Bus, error) {
	options := bus.NewOptions(opts...)

	if options.Embedded {
		return newEmbedded(nsqOptions(Options{}))
	}

	return nil, nil
}

func newEmbedded(opts *nsqd.Options) (*nsqEmbedded, error) {
	b, err := nsqd.New(opts)
	if err != nil {
		return nil, err
	}

	return &nsqEmbedded{
		opts: opts,
		nsqd: b,
		stop: make(chan struct{}),
	}, nil
}

func (nb *nsqEmbedded) Start() error {
	if err := nb.nsqd.Main(); err != nil {
		return err
	}

	<-nb.stop

	return nil
}

func (nb *nsqEmbedded) Shutdown() error {
	close(nb.stop)
	return nil
}

func (nb *nsqEmbedded) Mount(mod strana.Module) error {
	return mod.Events(nb)
}

func (nb *nsqEmbedded) Handle(src, sink message.Path, h strana.EventHandlerFunc) error {
	return nil
}

type nsqPublisher struct {
	p *nsq.Producer
}

func (nb *nsqEmbedded) Publisher() strana.Publisher {
	cfg := nsq.NewConfig()

	p, err := nsq.NewProducer(nb.opts.TCPAddress, cfg)
	if err != nil {
		return nil
	}

	return &nsqPublisher{
		p: p,
	}
}

func (np *nsqPublisher) Publish(topic string, m *message.Message) error {
	e, err := m.Envelope()
	if err != nil {
		return err
	}

	return np.p.Publish(topic, e)
}

func (np *nsqPublisher) Close() error {
	np.p.Stop()
	return nil
}

type nsqSubscriber struct {
	c    *nsq.Consumer
	opts *nsqd.Options
	h    strana.SubscriptionHandlerFunc
}

func (nb *nsqEmbedded) Subscriber() strana.Subscriber {
	return &nsqSubscriber{opts: nb.opts}
}

func (ns *nsqSubscriber) Subscribe(topic string, fn strana.SubscriptionHandlerFunc) error {
	ns.h = fn
	cfg := nsq.NewConfig()

	c, err := nsq.NewConsumer(topic, "subscribe", cfg)
	if err != nil {
		return err
	}

	ns.c = c

	c.AddHandler(nsq.HandlerFunc(ns.handle))

	return c.ConnectToNSQD(ns.opts.TCPAddress)
}

func (ns *nsqSubscriber) handle(m *nsq.Message) error {
	msg, err := message.Envelope(m.Body).Unmarshal()
	if err != nil {
		return err
	}

	return ns.h(msg)
}

func (ns *nsqSubscriber) Close() error {
	ns.c.Stop()
	return nil
}
