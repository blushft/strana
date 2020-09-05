package strana

import (
	"github.com/blushft/strana/event"
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

type ModuleConstructor func(config.Module) (Module, error)

type EventHandlerFunc func(*message.Message) ([]*message.Message, error)

type EventHandler interface {
	Handle(src message.Path, sink message.Path, hndlr EventHandlerFunc) error
	Publisher() Publisher
	Subscriber() Subscriber
}

type Publisher interface {
	Publish(topic string, e *message.Message) error
	Close() error
}

type SubscriptionHandlerFunc func(*message.Message) error

type Subscriber interface {
	Subscribe(topic string, fn SubscriptionHandlerFunc) error
	Close() error
}

type Producer interface {
	Subscribe(SubscriptionHandlerFunc) error
}

type Consumer interface {
	Publish(*event.Event) error
}

type Processor interface {
	Module
	Producer
	Consumer
}

type Source interface {
	Module
	Producer
}

type Sink interface {
	Module
	Consumer
}
