package loader

import (
	"log"

	"github.com/blushft/strana"
	"github.com/blushft/strana/pkg/event"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/store"
	"github.com/gofiber/fiber"
)

type Loader interface {
	strana.Module
	strana.Consumer
}

type Options struct{}

type loader struct {
	conf  config.Module
	opts  Options
	store *store.Store
	pub   strana.Publisher
}

func New(conf config.Module) (Loader, error) {
	return &loader{
		conf: conf,
	}, nil
}

func (l *loader) Routes(rtr fiber.Router) {}

func (l *loader) Events(eh strana.EventHandler) error {
	return eh.Subscriber().Subscribe(l.conf.Source.Topic, l.handle)
	return nil
}

func (l *loader) Services(s *store.Store) {

	l.store = s

}

func (l *loader) Publisher() strana.Publisher {
	return l.pub
}

func (l *loader) handle(msg *message.Message) error {

	evtType := event.Type(msg.Event.Event)
	switch evtType {
	default:
		log.Printf("unknown eventtype %s", evtType)
	}

	return nil
}

func (l *loader) storeMessage(evt *event.Event) error {
	panic("not implemented")
}
