package bus

import (
	"github.com/blushft/strana"
	"github.com/blushft/strana/platform/bus/message"
)

type Route struct {
	src message.Path
	pub strana.Publisher

	sink message.Path
	sub  strana.Subscriber

	hndlr strana.EventHandlerFunc
}

func NewRoute(src, sink message.Path, bus Bus, hndlr strana.EventHandlerFunc) (*Route, error) {
	r := &Route{
		src:   src,
		sink:  sink,
		hndlr: hndlr,
	}

	return r, nil
}

/* func (r *route) handleSub(msg *nats.Msg) {
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
} */

func (r *Route) Close() error {

	return nil
}
