package app

import (
	"log"
	"strconv"
	"strings"

	"github.com/blushft/strana/domain/entity"
	"github.com/blushft/strana/platform/store"
	"github.com/gofiber/fiber"
)

type reporter struct {
	pageviews entity.PageviewReporter
	sites     entity.AppReporter
}

func newReporter(api fiber.Router, s *store.Store) *reporter {
	pv := entity.NewPageviewService(s)
	sm := entity.NewAppService(s)

	r := &reporter{
		pageviews: pv,
		sites:     sm,
	}

	api.Get("/sites", r.listSites)
	api.Get("/pageviews", r.listPageviews)

	return r
}

func (r *reporter) listPageviews(c *fiber.Ctx) {
	params := c.Locals("params").(*Params)
	siteQ := c.Query("sites", "1")
	sites := []int{}

	if len(siteQ) > 0 {
		sqs := strings.Split(siteQ, ",")
		if len(sqs) > 0 {
			for _, sq := range sqs {
				sid, err := strconv.Atoi(sq)
				if err != nil {
					c.Status(400).Send("invalid site in query: " + sq)
					return
				}

				sites = append(sites, sid)
			}
		}
	}

	pvs, err := r.pageviews.List(params.Limit, params.Offset, sites...)
	if err != nil {
		log.Println("error fetching pageviews")
		c.Status(500).Send(err)
		return
	}

	if err := c.JSON(pvs); err != nil {
		log.Println("error marshalling json")
		c.Status(500).Send(err)
		return
	}
}

func (r *reporter) listSites(c *fiber.Ctx) {
	sites, err := r.sites.List()
	if err != nil {
		c.Status(500).Send(err)
	}

	if err := c.JSON(sites); err != nil {
		c.Status(500).Send(err)
	}
}
