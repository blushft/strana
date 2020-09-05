package collector

import (
	"encoding/base64"
	"errors"
	"log"
	"net/http"

	"github.com/blushft/strana"
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/platform"
	"github.com/blushft/strana/platform/bus/message"
	"github.com/blushft/strana/platform/cache"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
	"github.com/blushft/strana/platform/store"
	"github.com/blushft/strana/processor"
	"github.com/gofiber/fiber"
)

var ErrInvalidPayload = errors.New("could not collect event, invalid payload")

const (
	emptyGif = "R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7"
)

type TrackingCollector struct {
	conf       config.Module
	opts       Options
	log        *logger.Logger
	cache      *cache.Cache
	subscriber strana.Subscriber
	publisher  strana.Publisher

	procs []processor.EventProcessor
}

func newTrackingCollector(conf config.Module, opts Options) (*TrackingCollector, error) {
	c, err := cache.NewCache(opts.Cache)
	if err != nil {
		return nil, err
	}

	procs, err := platform.NewEventProcessorSet(opts.Processors)
	if err != nil {
		return nil, err
	}

	return &TrackingCollector{
		conf:  conf,
		cache: c,
		opts:  opts,
		procs: procs,
	}, nil
}

func (c *TrackingCollector) Routes(rtr fiber.Router) error {
	grp := rtr.Group("/analytics")

	grp.Get("/collect", c.collect)
	grp.Post("/collect", c.collect)

	return nil
}

func (c *TrackingCollector) Events(eh strana.EventHandler) error {
	c.publisher = eh.Publisher()
	c.subscriber = eh.Subscriber()

	return nil
}

func (c *TrackingCollector) Services(s *store.SQLStore) error {
	return nil
}

func (c *TrackingCollector) Logger(l *logger.Logger) {
	c.log = l.WithFields(logger.Fields{
		"module": "tracking_collector",
	})
}

func (c *TrackingCollector) Subscribe(fn strana.SubscriptionHandlerFunc) error {
	return c.subscriber.Subscribe(c.conf.Sink.Topic, fn)
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
		m := message.NewMessage(ne)

		if err := c.publisher.Publish(c.conf.Sink.Topic, m); err != nil {
			log.Printf("error publishing event: %v", err)
		}
	}
}

func (c *TrackingCollector) process(rm *event.Event) ([]*event.Event, error) {
	msgs, err := processor.Execute(c.procs, rm)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
