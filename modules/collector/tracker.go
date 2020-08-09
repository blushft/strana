package collector

import (
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/blushft/strana"
	"github.com/blushft/strana/domain/entity"
	"github.com/blushft/strana/platform/cache"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/store"
	"github.com/gofiber/fiber"
)

var ErrInvalidPayload = errors.New("could not collect event, invalid payload")

const (
	emptyGif = "R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7"
)

type TrackingCollector struct {
	conf      config.Module
	opts      Options
	cache     *cache.Cache
	sessions  entity.SessionManager
	publisher message.Publisher
}

func newTrackingCollector(conf config.Module, opts Options) (*TrackingCollector, error) {
	c, err := cache.NewCache(opts.Cache)
	if err != nil {
		return nil, err
	}

	return &TrackingCollector{
		conf:  conf,
		cache: c,
		opts:  opts,
	}, nil
}

func (c *TrackingCollector) Routes(rtr fiber.Router) {
	grp := rtr.Group("/analytics")

	grp.Get("/collect", c.collect)
	grp.Post("/collect", c.collect)
}

func (c *TrackingCollector) Events(eh strana.EventHandler) error {
	pb, err := eh.Broker(c.conf.Sink.Broker)
	if err != nil {
		return err
	}

	c.publisher = pb.Publisher()
	eh.Register(c.conf.Sink, c)

	return nil
}

func (c *TrackingCollector) Services(s *store.Store) {
	if c.cache == nil {
		c.sessions = entity.NewSessionService(s)
	} else {
		c.sessions = entity.NewCachedSessionService(s, c.cache)
	}
}

func (c *TrackingCollector) Publisher() message.Publisher {
	return c.publisher
}

func (c *TrackingCollector) collect(ctx *fiber.Ctx) {
	var rm entity.RawMessage

	switch ctx.Method() {
	case "POST":
		if err := ctx.BodyParser(&rm); err != nil {
			ctx.SendStatus(400)
			return
		}
	default:
		if err := ctx.QueryParser(&rm); err != nil {
			log.Printf("error parsing query: %v", err)
			ctx.SendStatus(400)
			return
		}
	}

	rm.IPAddress = "81.2.69.142"
	rm.UserAgent = string(ctx.Fasthttp.UserAgent())
	rm.Timestamp = time.Now().UTC().String()

	go c.publish(&rm)

	ctx.Set("Content-Type", "image/gif")
	ctx.Set("Expires", "Mon, 01 Jan 1990 00:00:00 GMT")
	ctx.Set("Cache-Control", "no-store")
	ctx.Set("Pragma", "no-cache")

	b, _ := base64.StdEncoding.DecodeString(emptyGif)
	ctx.Status(http.StatusOK).SendBytes(b)
}

func (c *TrackingCollector) publish(m *entity.RawMessage) {
	mb, err := m.JSON()
	if err != nil {
		log.Printf("error marshaling raw message: %v", err)
		return
	}

	msg := message.NewMessage(watermill.NewULID(), mb)

	if err := c.publisher.Publish(c.conf.Sink.Topic, msg); err != nil {
		log.Printf("error publishing event: %v", err)
	}
}