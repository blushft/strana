package bus

import (
	"github.com/blushft/strana"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/logger"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
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

func (nb *natsBus) newConn() (*nats.Conn, error) {
	return nats.Connect(nb.svr.ClientURL(), nats.Token(nb.token))
}

type publisher struct {
	conn *nats.Conn
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

func (np *publisher) Publish(topic string, e message.Envelope) error {
	return np.conn.Publish(topic, e)
}

func (np *publisher) Close() error {
	np.conn.Close()
	return nil
}

type subscriber struct {
	conn *nats.Conn
	sub  *nats.Subscription
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

func (ns *subscriber) Subscribe(topic string, fn func(*message.Message) error) error {
	ns.fn = fn
	sub, err := ns.conn.Subscribe(topic, ns.handle)
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
