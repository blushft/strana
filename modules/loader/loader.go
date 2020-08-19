package loader

import (
	"github.com/blushft/strana"
	"github.com/blushft/strana/pkg/event"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
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
	log   *logger.Logger
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
}

func (l *loader) Services(s *store.Store) {
	l.store = s
}

func (l *loader) Logger(lg *logger.Logger) {
	l.log = lg.WithFields(logger.Fields{"module": "loader"})
}

func (l *loader) Publisher() strana.Publisher {
	return l.pub
}

func (l *loader) handle(msg *message.Message) error {

	evtType := event.Type(msg.Event.Event)
	switch evtType {
	default:
		l.log.Warnf("unknown eventtype %s", evtType)
	}

	return nil
}

func (l *loader) storeMessage(evt *event.Event) error {
	panic("not implemented")
}
