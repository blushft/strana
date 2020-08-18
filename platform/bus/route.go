package bus

import (
	"log"

	"github.com/blushft/strana"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/nats-io/nats.go"
)

type route struct {
	source string
	sink   string
	conn   *nats.Conn
	sub    *nats.Subscription

	hndlr strana.EventHandlerFunc
}

func newRoute(src, sink string, conn *nats.Conn, hndlr strana.EventHandlerFunc) (*route, error) {
	r := &route{
		source: src,
		sink:   sink,
		conn:   conn,
		hndlr:  hndlr,
	}

	sub, err := r.conn.Subscribe(r.source, r.handleSub)
	if err != nil {
		return nil, err
	}

	r.sub = sub

	return r, nil
}

func (r *route) handleSub(msg *nats.Msg) {
	e := message.Envelope(msg.Data)
	m, err := e.Unmarshal()
	if err != nil {
		log.Printf("unable to unmarshal event envelope: %v", err)
	}

	msgs, err := r.hndlr(m)
	if err != nil {
		log.Printf("error handling event for subscription: %s, %v", r.source, err)
		return
	}

	for _, om := range msgs {
		oe, err := om.Envelope()
		if err != nil {
			log.Printf("error marshaling event enveleope: %v", err)
			continue
		}

		if err := r.conn.Publish(r.sink, oe); err != nil {
			log.Printf("error publishing result to %s: %v", r.sink, err)
			continue
		}
	}
}

func (r *route) Close() error {
	defer r.conn.Close()
	if err := r.sub.Unsubscribe(); err != nil {
		return err
	}

	return nil
}
