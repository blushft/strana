package reporter

import (
	"github.com/blushft/strana"
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/modules"
	"github.com/blushft/strana/modules/sink/reporter/command"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
	"github.com/blushft/strana/platform/store"
	"github.com/gofiber/fiber"
)

func init() {
	modules.Register("reporter", New)
}

type Reporter interface {
	strana.Sink
}

type reporter struct {
	conf config.Module
	log  *logger.Logger
	pub  strana.Publisher

	live *liveReporter

	evtExtractor *command.EventExtractor
	evtReporter  *command.EventReporter
}

func New(conf config.Module) (strana.Module, error) {
	return &reporter{
		conf: conf,
		live: newLiveReporter(),
	}, nil
}

func (mod *reporter) Routes(rtr fiber.Router) error {
	mod.routes(rtr)
	return nil
}

func (mod *reporter) Events(eh strana.EventHandler) error {
	go mod.live.run()

	return eh.Subscriber().Subscribe(mod.conf.Source, mod.handleEvents)
}

func (mod *reporter) Services(s *store.SQLStore) error {
	ee, err := command.NewEventExtractor(s)
	if err != nil {
		return err
	}

	er, err := command.NewEventReporter(s)
	if err != nil {
		return err
	}

	mod.evtExtractor = ee
	mod.evtReporter = er

	return nil
}

func (mod *reporter) Logger(lg *logger.Logger) {
	mod.log = lg.WithFields(logger.Fields{"module": "reporter"})
	mod.live.log = mod.log.WithFields(logger.Fields{"reporter": "live"})
}

func (mod *reporter) Publish(evt *event.Event) error {
	return mod.pub.Publish(mod.conf.Source, message.NewMessage(evt))
}

func (mod *reporter) handleEvents(msg *message.Message) error {
	if err := mod.evtExtractor.Save(msg.Event); err != nil {
		return err
	}

	mod.live.Send(string(msg.Event.Event), msg.Event)

	return nil
}
