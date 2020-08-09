package bus

import (
	"context"
	"errors"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/blushft/strana"
	"github.com/blushft/strana/platform/config"
)

type Bus struct {
	conf config.Bus
	log  watermill.LoggerAdapter
	rtr  *message.Router

	brokers map[string]PubSub
	sources map[string]Source
}

func New(conf config.Bus) (*Bus, error) {
	wla := watermill.NewStdLogger(conf.Debug, conf.Trace)

	rtr, err := message.NewRouter(message.RouterConfig{}, wla)
	if err != nil {
		return nil, err
	}

	brokers := make(map[string]PubSub, len(conf.Brokers))
	for n, b := range conf.Brokers {
		pb, err := NewPubSub(b, wla)
		if err != nil {
			return nil, err
		}

		brokers[n] = pb

	}

	return &Bus{
		conf:    conf,
		log:     wla,
		rtr:     rtr,
		brokers: brokers,
		sources: make(map[string]Source),
	}, nil
}

func (b *Bus) Register(src config.MessagePath, p strana.Producer) {
	b.sources[src.Topic] = Source{MessagePath: src, Producer: p}
}

func (b *Bus) Source(s string) (string, strana.Consumer, error) {
	src, ok := b.sources[s]
	if !ok {
		return "", nil, errors.New("no source found for " + s)
	}

	br, err := b.Broker(src.Broker)
	if err != nil {
		return "", nil, err
	}

	return src.Topic, br, nil
}

func (b *Bus) Mount(mod strana.Module) error {
	return mod.Events(b)
}

func (b *Bus) Router() *message.Router {
	return b.rtr
}

func (b *Bus) Broker(s string) (strana.Broker, error) {
	pb, ok := b.brokers[s]
	if !ok {
		return nil, errors.New("no broker found for " + s)
	}

	return pb, nil
}

func (b *Bus) Start() error {
	return b.rtr.Run(context.TODO())
}

func (b *Bus) Shutdown() error {
	return b.rtr.Close()
}
