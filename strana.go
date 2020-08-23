package strana

import (
	"github.com/blushft/strana/pkg/event"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
	"github.com/blushft/strana/platform/store"
	"github.com/gofiber/fiber"
)

type Module interface {
	Routes(fiber.Router) error
	Events(EventHandler) error
	Services(*store.SQLStore) error
	Logger(*logger.Logger)
}

type EventHandlerFunc func(*message.Message) ([]*message.Message, error)

type EventHandler interface {
	Handle(src string, sink string, hndlr EventHandlerFunc) error
	Producer
	Consumer
}

type Publisher interface {
	Publish(topic string, e message.Envelope) error
	Close() error
}

type Subscriber interface {
	Subscribe(topic string, fn func(*message.Message) error) error
	Close() error
}

type Producer interface {
	Subscriber() Subscriber
}

type Consumer interface {
	Publisher() Publisher
}

type Processor interface {
	Process(*event.Event) ([]*event.Event, error)
}

type ProcessorConstructor func(conf config.Processor) (Processor, error)

type Broker interface {
	Producer
	Consumer
}
