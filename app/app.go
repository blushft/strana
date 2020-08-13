package app

import (
	"log"

	"github.com/blushft/strana"
	"github.com/blushft/strana/modules"
	"github.com/blushft/strana/platform/bus"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/http"
	"github.com/blushft/strana/platform/store"
	"github.com/gofiber/fiber"
	"github.com/oklog/run"
)

type App struct {
	conf  config.Config
	svr   *http.Server
	bus   *bus.Bus
	store *store.Store

	modules map[string]strana.Module
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

	api := svr.Router().Group("/api", apiParams)

	api.Get("/config", func(c *fiber.Ctx) {
		if err := c.JSON(conf); err != nil {
			c.Status(500).Send(err)
		}
	})

	a := &App{
		conf:  conf,
		svr:   svr,
		bus:   bus,
		store: s,
	}

	return a, nil
}

func (a *App) initModules() error {
	mods := make(map[string]strana.Module, len(a.conf.Modules))

	for _, mconf := range a.conf.Modules {
		mod, err := modules.New(mconf)
		if err != nil {
			return err
		}

		a.svr.Mount(mod.Routes)
		if err := a.bus.Mount(mod); err != nil {
			return err
		}

		a.store.Mount(mod.Services)

		mods[mconf.Name] = mod
	}

	a.modules = mods

	return nil
}

func (a *App) Start() error {
	if err := a.initModules(); err != nil {
		return err
	}

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
