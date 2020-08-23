package loader

import (
	"github.com/blushft/strana"
	"github.com/blushft/strana/domain/loader/entity"
	"github.com/blushft/strana/pkg/event"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
	"github.com/blushft/strana/platform/store"
	ls "github.com/blushft/strana/platform/store/loader"
	"github.com/gofiber/fiber"
)

type Loader interface {
	strana.Module
	strana.Consumer
}

type Options struct{}

type loader struct {
	conf config.Module
	opts Options
	log  *logger.Logger
	pub  strana.Publisher

	restore entity.RawEventManager
}

func New(conf config.Module) (Loader, error) {
	return &loader{
		conf: conf,
	}, nil
}

func (l *loader) Routes(rtr fiber.Router) error {
	return nil
}

func (l *loader) Events(eh strana.EventHandler) error {
	return eh.Subscriber().Subscribe(l.conf.Source.Topic, l.handle)
}

func (l *loader) Services(s *store.SQLStore) error {
	dbc, err := ls.NewStore(s)
	if err != nil {
		return err
	}

	l.restore = entity.NewRawEventService(dbc)

	return nil
}

func (l *loader) Logger(lg *logger.Logger) {
	l.log = lg.WithFields(logger.Fields{"module": "loader"})
}

func (l *loader) Publisher() strana.Publisher {
	return l.pub
}

func (l *loader) handle(msg *message.Message) error {
	return l.storeMessage(msg.Event)
}

func (l *loader) storeMessage(evt *event.Event) error {
	re := &entity.RawEvent{Event: evt}

	return l.restore.Create(re)
}
