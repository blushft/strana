package loader

import (
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/blushft/strana"
	"github.com/blushft/strana/domain/command"
	"github.com/blushft/strana/domain/entity"
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
	sub   message.Subscriber

	pvcmd command.TrackPageviewCommand
}

func New(conf config.Module) (Loader, error) {
	return &loader{
		conf: conf,
	}, nil
}

func (l *loader) Routes(rtr fiber.Router) {}

func (l *loader) Events(eh strana.EventHandler) error {
	topic, sub, err := eh.Source(l.conf.Source.Topic)
	if err != nil {
		return err
	}

	l.sub = sub.Subscriber()

	eh.Router().AddNoPublisherHandler(
		"loader",
		topic,
		l.sub,
		l.handle,
	)

	return nil
}

func (l *loader) Services(s *store.Store) {
	appmgr := entity.NewAppService(s)
	sesmgr := entity.NewSessionService(s)
	pvmgr := entity.NewPageviewService(s)
	usrmgr := entity.NewUserService(s)

	l.store = s
	l.pvcmd = command.NewTrackPageviewCommand(appmgr, pvmgr, sesmgr, usrmgr)
}

func (l *loader) Subscriber() message.Subscriber {
	return l.sub
}

func (l *loader) handle(msg *message.Message) error {
	rm, err := entity.RawMessageFromPayload(msg)
	if err != nil {
		return err
	}

	evt := entity.EventType(rm.Event)
	switch evt {
	case entity.EventTypePageview:
		return l.pvcmd.Track(rm)
	default:
		log.Printf("unknown eventtype %s", evt)
	}

	return nil
}

func (l *loader) storeMessage(msg *entity.RawMessage) error {
	panic("not implemented")
}
