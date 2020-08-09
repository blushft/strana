package enhancer

import (
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/blushft/strana"
	"github.com/blushft/strana/domain/entity"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/store"
	"github.com/blushft/strana/processors"
	"github.com/gofiber/fiber"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

type Enhancer interface {
	strana.Module
	strana.Consumer
	strana.Producer
}

type Options struct {
	Processors []config.Processor `json:"processors" yaml:"processors" mapstructure:"processors"`
}

type enhancer struct {
	conf config.Module
	opts Options

	pub message.Publisher
	sub message.Subscriber

	procs []strana.Processor
}

func New(conf config.Module) (Enhancer, error) {
	opts := Options{}
	if err := mapstructure.Decode(conf.Options, &opts); err != nil {
		return nil, err
	}

	procs := make([]strana.Processor, 0, len(opts.Processors))

	for _, p := range opts.Processors {
		proc, err := processors.New(p)
		if err != nil {
			return nil, err
		}

		procs = append(procs, proc)
	}

	return &enhancer{
		conf:  conf,
		opts:  opts,
		procs: procs,
	}, nil
}

func (e *enhancer) Routes(fiber.Router) {}

func (e *enhancer) Services(*store.Store) {}

func (e *enhancer) Events(eh strana.EventHandler) error {
	topic, sub, err := eh.Source(e.conf.Source.Topic)
	if err != nil {
		return err
	}

	pb, err := eh.Broker(e.conf.Sink.Broker)
	if err != nil {
		return err
	}

	e.sub = sub.Subscriber()
	e.pub = pb.Publisher()

	eh.Register(e.conf.Sink, e)

	eh.Router().AddHandler(
		e.conf.Name,
		topic,
		e.sub,
		e.conf.Sink.Topic,
		e.pub,
		e.handle,
	)

	return nil
}

func (e *enhancer) Publisher() message.Publisher {
	return e.pub
}

func (e *enhancer) Subscriber() message.Subscriber {
	return e.sub
}

func (e *enhancer) handle(msg *message.Message) ([]*message.Message, error) {
	if msg.Payload == nil {
		log.Println("message is nil bro")
		return nil, nil
	}

	var rm *entity.RawMessage
	if err := json.Unmarshal(msg.Payload, &rm); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal payload: "+e.conf.Name)
	}

	q := []*entity.RawMessage{rm}

	for i := 0; len(q) > 0 && i < len(e.procs); i++ {
		var nextQ []*entity.RawMessage
		for _, m := range q {
			res, err := e.procs[i].Process(m)
			if err != nil {
				return nil, err
			}

			nextQ = append(nextQ, res...)
		}

		q = nextQ
	}

	if len(q) == 0 {
		return nil, nil
	}

	results := make([]*message.Message, 0, len(q))

	for _, qm := range q {
		pl, err := qm.JSON()
		if err != nil {
			return nil, err
		}

		md := msg.Copy().Metadata

		nm := message.NewMessage(watermill.NewULID(), pl)
		nm.Metadata = md
		results = append(results, nm)
	}

	return results, nil
}
