package strana

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/blushft/strana/pkg/event"
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
	Register(config.MessagePath, Producer)
	Source(string) (string, Consumer, error)
}

type Producer interface {
	Publisher() message.Publisher
}

type Consumer interface {
	Subscriber() message.Subscriber
}

type Processor interface {
	Process(*event.Event) ([]*event.Event, error)
}

type ProcessorConstructor func(conf config.Processor) (Processor, error)

type Broker interface {
	Producer
	Consumer
}
