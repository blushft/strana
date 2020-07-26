package bus

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/blushft/strana/platform"
	"github.com/blushft/strana/platform/config"
)

func errInvalidConfig(s string) error {
	return fmt.Errorf("invalid pubsub driver: %s", s)
}

type PubSub interface {
	platform.Producer
	platform.Consumer
}

func NewPubSub(conf config.PubSub, l watermill.LoggerAdapter) (PubSub, error) {
	switch conf.Driver {
	case "memory":

		ps := NewMemoryPubSub(conf, l)
		return &pubsub{
			pub: ps,
			sub: ps,
		}, nil
	default:
		return nil, errInvalidConfig(conf.Driver)
	}
}

type pubsub struct {
	pub message.Publisher
	sub message.Subscriber
}

func (pb *pubsub) Publisher() message.Publisher {
	return pb.pub
}

func (pb *pubsub) Subscriber() message.Subscriber {
	return pb.sub
}

func NewMemoryPubSub(conf config.PubSub, l watermill.LoggerAdapter) *gochannel.GoChannel {
	goconf := &gochannel.Config{}
	_ = conf.DriverConfig(goconf)

	return gochannel.NewGoChannel(*goconf, l)
}

type Source struct {
	config.Source
	platform.Producer
}

type Sink struct {
	config.Sink
	platform.Consumer
}
