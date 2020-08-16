package bus

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-nats/pkg/nats"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/blushft/strana"
	"github.com/blushft/strana/platform/config"
	"github.com/mitchellh/mapstructure"
	"github.com/nats-io/stan.go"
)

func errInvalidConfig(s string) error {
	return fmt.Errorf("invalid pubsub driver: %s", s)
}

type Source struct {
	config.MessagePath
	strana.Producer
}

type Sink struct {
	config.MessagePath
	strana.Consumer
}

type PubSub interface {
	strana.Producer
	strana.Consumer
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

type natsConfig struct {
	URL         string `json:"url" mapstructure:"url"`
	ClusterID   string `json:"cluster_id" mapstructure:"cluster_id"`
	ClientID    string `json:"client_id" mapstructure:"client_id"`
	QueueGroup  string `json:"queue_group" mapstructure:"queue_group"`
	DurableName string `json:"durable_name" mapstructure:"durable_name"`
}

func NewNatsPubSub(conf config.PubSub, l watermill.LoggerAdapter) (*pubsub, error) {
	var nconf natsConfig
	if err := mapstructure.Decode(conf.Config, &nconf); err != nil {
		return nil, err
	}

	stanOpts := []stan.Option{
		stan.NatsURL(nconf.URL),
	}

	codec := nats.GobMarshaler{}

	subConf := nats.StreamingSubscriberConfig{
		ClusterID:   nconf.ClusterID,
		ClientID:    nconf.ClientID,
		QueueGroup:  nconf.QueueGroup,
		DurableName: nconf.DurableName,
		StanOptions: stanOpts,
		Unmarshaler: codec,
	}

	sub, err := nats.NewStreamingSubscriber(subConf, l)
	if err != nil {
		return nil, err
	}

	pubConf := nats.StreamingPublisherConfig{
		ClusterID:   nconf.ClusterID,
		ClientID:    nconf.ClientID,
		StanOptions: stanOpts,
		Marshaler:   codec,
	}

	pub, err := nats.NewStreamingPublisher(pubConf, l)
	if err != nil {
		return nil, err
	}

	return &pubsub{
		pub: pub,
		sub: sub,
	}, nil
}
