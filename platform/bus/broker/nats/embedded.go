package nats

import (
	"github.com/blushft/strana"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/logger"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	natsc "github.com/nats-io/nats.go"
)

type natsBus struct {
	opts  server.Options
	token string

	svr *server.Server
}

func newNatsBus(opts server.Options) (*natsBus, error) {
	svr, err := server.NewServer(&opts)
	if err != nil {
		return nil, err
	}

	return &natsBus{
		opts:  opts,
		token: opts.Authorization,
		svr:   svr,
	}, nil
}

func (nb *natsBus) Start() {
	nb.svr.Start()
	nb.svr.WaitForShutdown()
}

func (nb *natsBus) Shutdown() {
	nb.svr.Shutdown()
}

func (nb *natsBus) newConn() (*natsc.Conn, error) {
	return natsc.Connect(nb.svr.ClientURL(), natsc.Token(nb.token))
}

type publisher struct {
	conn *natsc.Conn
}

func (nb *natsBus) NewPublisher() (strana.Publisher, error) {
	conn, err := nb.newConn()
	if err != nil {
		return nil, err
	}

	return &publisher{
		conn: conn,
	}, nil
}

func (np *publisher) Publish(p message.Path, m *message.Message) error {
	e, err := m.Envelope()
	if err != nil {
		return err
	}

	return np.conn.Publish(p.Topic, e)
}

func (np *publisher) Close() error {
	np.conn.Close()
	return nil
}

type subscriber struct {
	conn *natsc.Conn
	sub  *natsc.Subscription
	fn   func(*message.Message) error
}

func (nb *natsBus) NewSubscriber() (strana.Subscriber, error) {
	conn, err := nb.newConn()
	if err != nil {
		return nil, err
	}

	s := &subscriber{
		conn: conn,
	}

	return s, nil
}

func (ns *subscriber) Subscribe(p message.Path, fn strana.SubscriptionHandlerFunc) error {
	ns.fn = fn
	sub, err := ns.conn.Subscribe(p.Topic, ns.handle)
	if err != nil {
		return err
	}

	ns.sub = sub

	return nil
}

func (ns *subscriber) Close() error {
	defer ns.conn.Close()
	if err := ns.sub.Unsubscribe(); err != nil {
		return err
	}

	return nil
}

func (ns *subscriber) handle(msg *nats.Msg) {
	m, err := message.Envelope(msg.Data).Unmarshal()
	if err != nil {
		logger.Log().WithError(err).Error("error extracting bus message")
		return
	}

	if err := ns.fn(m); err != nil {
		logger.Log().WithError(err).Error("error handling bus message")
	}
}
