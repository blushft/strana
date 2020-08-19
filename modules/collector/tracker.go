package collector

import (
	"encoding/base64"
	"errors"
	"log"
	"net/http"

	"github.com/blushft/strana"
	"github.com/blushft/strana/domain/entity"
	"github.com/blushft/strana/pkg/event"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/cache"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
	"github.com/blushft/strana/platform/store"
	"github.com/blushft/strana/processors"
	"github.com/gofiber/fiber"
)

var ErrInvalidPayload = errors.New("could not collect event, invalid payload")

const (
	emptyGif = "R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7"
)

type TrackingCollector struct {
	conf      config.Module
	opts      Options
	log       *logger.Logger
	cache     *cache.Cache
	sessions  entity.SessionManager
	publisher strana.Publisher

	procs []strana.Processor
}

func newTrackingCollector(conf config.Module, opts Options) (*TrackingCollector, error) {
	c, err := cache.NewCache(opts.Cache)
	if err != nil {
		return nil, err
	}

	procs := make([]strana.Processor, 0, len(opts.Processors))
	for _, p := range opts.Processors {
		proc, err := processors.New(p)
		if err != nil {
			return nil, err
		}

		procs = append(procs, proc)
	}

	return &TrackingCollector{
		conf:  conf,
		cache: c,
		opts:  opts,
		procs: procs,
	}, nil
}

func (c *TrackingCollector) Routes(rtr fiber.Router) {
	grp := rtr.Group("/analytics")

	grp.Get("/collect", c.collect)
	grp.Post("/collect", c.collect)
}

func (c *TrackingCollector) Events(eh strana.EventHandler) error {
	c.publisher = eh.Publisher()

	return nil
}

func (c *TrackingCollector) Services(s *store.Store) {
	if c.cache == nil {
		c.sessions = entity.NewSessionService(s)
	} else {
		c.sessions = entity.NewCachedSessionService(s, c.cache)
	}
}

func (c *TrackingCollector) Logger(l *logger.Logger) {
	c.log = l.WithFields(logger.Fields{
		"module": "tracking_collector",
	})
}

func (c *TrackingCollector) Publisher() strana.Publisher {
	return c.publisher
}

func (c *TrackingCollector) collect(ctx *fiber.Ctx) {
	rm := event.Empty()

	switch ctx.Method() {
	case "POST":
		if err := ctx.BodyParser(rm); err != nil {
			log.Printf("error binding message: %v", err)
			ctx.SendStatus(400)
			return
		}
	default:
		if err := ctx.QueryParser(rm); err != nil {
			log.Printf("error parsing query: %v", err)
			ctx.SendStatus(400)
			return
		}
	}

	rm.SetContext(event.NewNetworkContext("24.106.166.33", string(ctx.Fasthttp.UserAgent())))

	go c.publish(rm)

	ctx.Set("Content-Type", "image/gif")
	ctx.Set("Expires", "Mon, 01 Jan 1990 00:00:00 GMT")
	ctx.Set("Cache-Control", "no-store")
	ctx.Set("Pragma", "no-cache")

	b, _ := base64.StdEncoding.DecodeString(emptyGif)
	ctx.Status(http.StatusOK).SendBytes(b)
}

func (c *TrackingCollector) publish(evt *event.Event) {
	evts, err := c.process(evt)
	if err != nil {
		log.Printf("error processing messages: %v", err)
		return
	}

	for _, ne := range evts {
		e, err := message.NewMessage(ne).Envelope()
		if err != nil {
			log.Printf("error marshaling raw message: %v", err)
			continue
		}

		if err := c.publisher.Publish(c.conf.Sink.Topic, e); err != nil {
			log.Printf("error publishing event: %v", err)
		}
	}
}

func (c *TrackingCollector) process(rm *event.Event) ([]*event.Event, error) {
	msgs, err := processors.Execute(c.procs, rm)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
