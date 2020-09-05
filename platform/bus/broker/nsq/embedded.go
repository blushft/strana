package nsq

import (
	"github.com/blushft/strana"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/logger"
	"github.com/nsqio/go-nsq"
	"github.com/nsqio/nsq/nsqadmin"
	"github.com/nsqio/nsq/nsqd"
	"github.com/oklog/run"
	"github.com/pkg/errors"
)

type nsqEmbedded struct {
	opts     Options
	nsqd     *nsqd.NSQD
	nsqadmin *nsqadmin.NSQAdmin

	stop chan struct{}
}

func newEmbedded(opts Options) (*nsqEmbedded, error) {
	b, err := nsqd.New(opts.NSQOptions)
	if err != nil {
		return nil, err
	}

	var admin *nsqadmin.NSQAdmin
	if opts.BrokerOptions.Web {
		aopts := nsqadmin.NewOptions()
		aopts.NSQDHTTPAddresses = []string{opts.NSQOptions.HTTPAddress}
		admin, err = nsqadmin.New(aopts)
		if err != nil {
			return nil, errors.Wrap(err, "new nsqadmin")
		}
	}

	return &nsqEmbedded{
		opts:     opts,
		nsqd:     b,
		nsqadmin: admin,
		stop:     make(chan struct{}),
	}, nil
}

func (nb *nsqEmbedded) Connect() error {
	g := run.Group{}

	g.Add(nb.nsqd.Main, func(error) {
		nb.nsqd.Exit()
		logger.Log().Info("nsqd exit")
	})

	if nb.nsqadmin != nil {
		g.Add(nb.nsqadmin.Main, func(error) {
			nb.nsqadmin.Exit()
			logger.Log().Info("nsqadmin exit")
		})
	}

	g.Add(func() error {
		<-nb.stop
		return nil
	}, func(error) {})

	return g.Run()
}

func (nb *nsqEmbedded) Disconnect() error {
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

	p, err := nsq.NewProducer(nb.opts.NSQOptions.TCPAddress, cfg)
	if err != nil {
		return nil
	}

	return &nsqPublisher{
		p: p,
	}
}

func (np *nsqPublisher) Publish(p message.Path, m *message.Message) error {
	e, err := m.Envelope()
	if err != nil {
		return err
	}

	return np.p.Publish(p.Topic, e)
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
	return &nsqSubscriber{opts: nb.opts.NSQOptions}
}

func (ns *nsqSubscriber) Subscribe(p message.Path, fn strana.SubscriptionHandlerFunc) error {
	ns.h = fn
	cfg := nsq.NewConfig()

	c, err := nsq.NewConsumer(p.Topic, p.Channel, cfg)
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
