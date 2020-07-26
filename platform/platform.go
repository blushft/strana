package platform

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/store"
	"github.com/gofiber/fiber"
)

type Module interface {
	Routes(fiber.Router)
	Events(EventHandler) error
	Services(*store.Store)
}

type EventHandler interface {
	Broker(string) (Broker, error)
	Router() *message.Router
	Register(config.Source, Producer)
	Source(string) (string, Consumer, error)
}

type Producer interface {
	Publisher() message.Publisher
}

type Consumer interface {
	Subscriber() message.Subscriber
}

type Broker interface {
	Producer
	Consumer
}
