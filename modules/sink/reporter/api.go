package reporter

import (
	"net/http"

	"github.com/blushft/strana/modules/sink/reporter/entity"
	"github.com/gofiber/fiber"
	"github.com/gofiber/websocket"
)

func (mod *reporter) routes(rtr fiber.Router) {
	api := rtr.Group("/reporter")

	mod.liveRoutes(api)
	mod.reportRoutes(api)
}

func (mod *reporter) liveRoutes(rtr fiber.Router) {
	grp := rtr.Group("/live")

	grp.Get(
		"/events",
		func(c *fiber.Ctx) {
			if websocket.IsWebSocketUpgrade(c) {
				c.Next()
			}
		},
		websocket.New(mod.live.handleLive),
	)

	grp.Get("/rates", mod.live.handleRates)
}

func (mod *reporter) reportRoutes(rtr fiber.Router) {
	rtr.Get("/events", getParams, mod.handleGetEvents)
	rtr.Get("/events/count", mod.handleGetEventsCount)
	rtr.Get("/events/actions", getParams, mod.handleGetEventsActions)
	rtr.Get("/events/actions/top", getParams, mod.handleGetTopActions)
}

func getParams(c *fiber.Ctx) {
	var params entity.QueryParams
	if err := c.QueryParser(params); err != nil {
		c.Status(http.StatusBadRequest).Send(err)

		return
	}

	c.Locals("params", params)

	c.Next()
}

func (mod *reporter) handleGetEvents(c *fiber.Ctx) {
	params := c.Locals("params").(entity.QueryParams)

	evts, err := mod.evtReporter.Events(params)
	if err != nil {
		c.Status(http.StatusInternalServerError).Send(err)

		return
	}

	if err := c.JSON(evts); err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
		return
	}
}

func (mod *reporter) handleGetEventsCount(c *fiber.Ctx) {
	count, err := mod.evtReporter.EventsCount()
	if err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
	}

	res := map[string]int{"count": count}

	if err := c.JSON(res); err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
	}
}

func (mod *reporter) handleGetEventsActions(c *fiber.Ctx) {
	params := c.Locals("params").(entity.QueryParams)

	evts, err := mod.evtReporter.EventsWithAction(params)
	if err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
		return
	}

	if err := c.JSON(evts); err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
		return
	}
}

func (mod *reporter) handleGetTopActions(c *fiber.Ctx) {
	params := c.Locals("params").(entity.QueryParams)

	cats, err := mod.evtReporter.TopActionCatgories(params)
	if err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
		return
	}

	if err := c.JSON(cats); err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
		return
	}
}
