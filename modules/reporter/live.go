package reporter

import (
	"time"

	"github.com/blushft/strana/event"
	"github.com/blushft/strana/platform/logger"
	"github.com/gofiber/fiber"
	"github.com/gofiber/websocket"
	"github.com/paulbellamy/ratecounter"
)

type liveReporter struct {
	liveClients map[*websocket.Conn]liveClient
	reg         chan *websocket.Conn
	dereg       chan *websocket.Conn
	live        chan map[string]*event.Event

	actionCount *ratecounter.RateCounter

	log *logger.Logger
}

func newLiveReporter() *liveReporter {
	return &liveReporter{
		liveClients: make(map[*websocket.Conn]liveClient),
		reg:         make(chan *websocket.Conn),
		dereg:       make(chan *websocket.Conn),
		live:        make(chan map[string]*event.Event, 1),
		actionCount: ratecounter.NewRateCounter(time.Minute * 30),
	}
}

type liveClient struct {
}

func (r *liveReporter) run() {
	for {
		select {
		case conn := <-r.reg:
			r.liveClients[conn] = liveClient{}
		case e := <-r.live:
			var cc *websocket.Conn
			for conn := range r.liveClients {
				cc = conn
				if err := cc.WriteJSON(e); err != nil {
					r.log.WithError(err).Error("error writing to client")
					r.dereg <- cc
					cc.Close()
				}
			}
		case conn := <-r.dereg:

			delete(r.liveClients, conn)
		}
	}
}

func (r *liveReporter) Send(channel string, evt *event.Event) {
	switch channel {
	case "action":
		r.actionCount.Incr(1)
	}

	r.live <- map[string]*event.Event{channel: evt}
}

func (r *liveReporter) handleRates(c *fiber.Ctx) {
	r.log.Debug("getting event rates")
	rates := map[string]int64{
		"actions": r.actionCount.Rate(),
	}

	if err := c.JSON(rates); err != nil {
		r.log.WithError(err).Error("error handling rates")
		c.Status(500).Send(err)
		return
	}
}

func (r *liveReporter) handleLive(c *websocket.Conn) {
	r.log.Debug("client connected to live stream")

	defer func() {
		r.dereg <- c
		c.Close()
	}()

	r.reg <- c

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				r.log.WithError(err).Error("read error")
			}

			return
		}

		r.log.WithFields(logger.Fields{"type": mt, "msg": msg}).Info("message received")
	}
}
