package nsq

import (
	"github.com/blushft/strana/platform/bus/broker"
)

func NewDefault(opts ...broker.Option) broker.Broker {
	b, err := New(opts...)
	if err != nil {
		panic(err)
	}

	return b
}

func New(opts ...broker.Option) (broker.Broker, error) {
	options := broker.NewOptions(opts...)

	if options.Embedded {
		return newEmbedded(newOptions(options))
	}

	return nil, nil
}
