package loader

import (
	"github.com/blushft/strana"

	"github.com/blushft/strana/event"
	"github.com/blushft/strana/modules"
	"github.com/blushft/strana/modules/sink/loader/entity"
	ls "github.com/blushft/strana/modules/sink/loader/store"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
	"github.com/blushft/strana/platform/store"
	"github.com/gofiber/fiber"
)

func init() {
	modules.Register("loader", New)
}

type Loader interface {
	strana.Sink
}

type loader struct {
	conf config.Module

	log *logger.Logger
	pub strana.Publisher

	restore entity.RawEventManager
}

func New(conf config.Module) (strana.Module, error) {
	return &loader{
		conf: conf,
	}, nil
}

func (l *loader) Routes(rtr fiber.Router) error {
	return nil
}

func (l *loader) Events(eh strana.EventHandler) error {
	l.pub = eh.Publisher()
	return eh.Subscriber().Subscribe(l.conf.Source, l.handle)
}

func (l *loader) Services(s *store.SQLStore) error {
	dbc, err := ls.New(s)
	if err != nil {
		return err
	}

	l.restore = entity.NewRawEventService(dbc)

	return nil
}

func (l *loader) Logger(lg *logger.Logger) {
	l.log = lg.WithFields(logger.Fields{"module": "loader"})
}

func (l *loader) Publish(evt *event.Event) error {
	return l.pub.Publish(l.conf.Source, message.NewMessage(evt))
}

func (l *loader) handle(msg *message.Message) error {
	return l.storeMessage(msg.Event)
}

func (l *loader) storeMessage(evt *event.Event) error {
	re := &entity.RawEvent{Event: evt}

	return l.restore.Create(re)
}
