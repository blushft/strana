package reporter

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
	"github.com/gofiber/websocket"
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

	restore entity.RawEventReporter
}

func New(conf config.Module) (strana.Module, error) {
	return &reporter{
		conf: conf,
		live: newLiveReporter(),
	}, nil
}

func (mod *reporter) Routes(rtr fiber.Router) error {
	api := rtr.Group("/reporter")

	live := api.Group("/live")

	live.Get(
		"/events",
		func(c *fiber.Ctx) {
			if websocket.IsWebSocketUpgrade(c) {
				c.Next()
			}
		},
		websocket.New(mod.live.handleLive),
	)

	live.Get("/rates", mod.live.handleRates)

	return nil
}

func (mod *reporter) Events(eh strana.EventHandler) error {
	go mod.live.run()

	return eh.Subscriber().Subscribe(mod.conf.Source.Topic, mod.handleEvents)
}

func (mod *reporter) Services(s *store.SQLStore) error {
	dbc, err := ls.New(s)
	if err != nil {
		return err
	}

	mod.restore = entity.NewRawEventService(dbc)

	return nil
}

func (mod *reporter) Logger(lg *logger.Logger) {
	mod.log = lg.WithFields(logger.Fields{"module": "reporter"})
	mod.live.log = mod.log.WithFields(logger.Fields{"reporter": "live"})
}

func (mod *reporter) Publish(evt *event.Event) error {
	return mod.pub.Publish(mod.conf.Source.Source, message.NewMessage(evt))
}

func (mod *reporter) handleEvents(msg *message.Message) error {
	mod.live.Send(string(msg.Event.Event), msg.Event)

	return nil
}
