package platform

import (
	"errors"

	"github.com/blushft/strana/platform/bus"
	"github.com/blushft/strana/platform/bus/nsq"
	"github.com/blushft/strana/platform/config"
)

var ErrInvalidBroker = errors.New("invalid broker")

func NewBus(conf config.Bus) (bus.Bus, error) {
	switch conf.Broker {
	case "nsq":
		return nsq.New()
	default:
		return nil, ErrInvalidBroker
	}
}
