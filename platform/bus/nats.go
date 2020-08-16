package bus

import (
	"encoding/json"
	"log"

	"github.com/blushft/strana/pkg/event"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

type natsBus struct {
	opts  server.Options
	token string

	svr *server.Server
}

func newNatsBus(busport, httpport int, token string) (*natsBus, error) {
	opts := server.Options{
		Port:          busport,
		HTTPPort:      httpport,
		Authorization: token,
	}

	svr, err := server.NewServer(&opts)
	if err != nil {
		return nil, err
	}

	return &natsBus{
		opts:  opts,
		token: token,
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

func (nb *natsBus) Publish(topic string, evt *event.Event) error {
	conn, err := nb.newConn()
	if err != nil {
		return err
	}

	b, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	return conn.Publish(topic, b)
}

func (nb *natsBus) NewSubscriber(topic string, hndlr subsciberFunc) (*subscriber, error) {
	conn, err := nb.newConn()
	if err != nil {
		return nil, err
	}

	sub := &subscriber{
		conn:  conn,
		hndlr: hndlr,
	}

	ns, err := conn.Subscribe(topic, sub.handle)
	if err != nil {
		return nil, err
	}

	sub.sub = ns

	return sub, nil
}

type subsciberFunc func(*event.Event) error

type subscriber struct {
	conn  *nats.Conn
	sub   *nats.Subscription
	hndlr subsciberFunc
}

func (s *subscriber) handle(msg *nats.Msg) {
	evt := event.Empty()
	if err := json.Unmarshal(msg.Data, &evt); err != nil {
		log.Printf("error getting event for subscription: %s - %v\n", s.sub.Subject, err)
	}

	if err := s.hndlr(evt); err != nil {
		log.Printf("error handling event for subscription: %s, %v", s.sub.Subject, err)
	}
}

func (s *subscriber) Close() error {
	defer s.conn.Close()
	if err := s.sub.Unsubscribe(); err != nil {
		return err
	}

	return nil
}
