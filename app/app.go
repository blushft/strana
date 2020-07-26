package app

import (
	"log"

	"github.com/blushft/strana/collector"
	"github.com/blushft/strana/enhancer"
	"github.com/blushft/strana/platform/bus"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/http"
	"github.com/blushft/strana/platform/store"
	"github.com/gofiber/fiber"
	"github.com/oklog/run"
	"github.com/pkg/errors"
)

type App struct {
	conf      config.Config
	svr       *http.Server
	bus       *bus.Bus
	store     *store.Store
	collector collector.Collector
	enhancer  enhancer.Enhancer
	reporter  *reporter
}

func New(conf config.Config) (*App, error) {
	s, err := store.NewStore(conf.Database)
	if err != nil {
		return nil, err
	}

	svr := http.NewServer(conf.Server, conf.Debug)

	bus, err := bus.New(conf.Bus)
	if err != nil {
		return nil, err
	}

	coll, err := collector.NewCollector(conf.Collector)
	if err != nil {
		return nil, errors.Wrap(err, "new collector")
	}

	svr.Mount(coll.Routes)
	if err := bus.Mount(coll); err != nil {
		return nil, errors.Wrap(err, "mount collector")
	}
	s.Mount(coll.Services)

	enh := enhancer.NewEnhancer(conf.Enhancer)

	if err := bus.Mount(enh); err != nil {
		return nil, errors.Wrap(err, "mount enhancer")
	}

	api := svr.Router().Group("/api", apiParams)

	api.Get("/config", func(c *fiber.Ctx) {
		if err := c.JSON(conf); err != nil {
			c.Status(500).Send(err)
		}
	})

	rptr := newReporter(api, s)

	a := &App{
		conf:      conf,
		svr:       svr,
		bus:       bus,
		store:     s,
		collector: coll,
		enhancer:  enh,
		reporter:  rptr,
	}

	return a, nil
}

func (a *App) Start() error {
	grp := run.Group{}

	grp.Add(
		a.svr.Start,
		func(e error) {
			log.Println("server stopping")
			a.svr.Shutdown()
		},
	)

	grp.Add(
		a.bus.Start,
		func(e error) {
			log.Println("bus stopping")
			a.bus.Shutdown()
		},
	)

	return grp.Run()
}
