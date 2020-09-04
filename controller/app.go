package controller

import (
	"github.com/blushft/strana"
	"github.com/blushft/strana/modules"
	"github.com/blushft/strana/platform"
	"github.com/blushft/strana/platform/bus"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
	"github.com/blushft/strana/platform/server"
	"github.com/blushft/strana/platform/store"
	"github.com/gofiber/fiber"
	"github.com/oklog/run"
)

type App struct {
	conf  config.Config
	svr   *server.Server
	bus   bus.Bus
	store *store.SQLStore
	log   *logger.Logger

	modules map[string]strana.Module
}

func New(conf config.Config) (*App, error) {
	logger.Init(conf.Logger)

	l := logger.New().WithFields(logger.Fields{"app": "strana"})

	s, err := store.NewSQL(conf.Database)
	if err != nil {
		return nil, err
	}

	svr := server.New(conf.Server, conf.Debug)

	bus, err := platform.NewBus(conf.Bus)
	if err != nil {
		return nil, err
	}

	api := svr.Router().Group("/api")

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
		log:   l,
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

		if err := a.svr.Mount(mod.Routes); err != nil {
			return err
		}

		if err := a.store.Mount(mod.Services); err != nil {
			return err
		}

		a.log.Mount(mod.Logger)

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
			a.log.Info("server stopping")
			a.svr.Shutdown()
		},
	)

	grp.Add(
		a.bus.Start,
		func(e error) {
			a.log.Info("bus stopping")
			if err := a.bus.Shutdown(); err != nil {
				a.log.WithError(err).Error("bus shutdown")
			}
		},
	)

	go func() {
		for k, mod := range a.modules {
			a.log.Infof("mounting events for module %s", k)

			if err := a.bus.Mount(mod); err != nil {
				logger.Log().Fatal(err.Error())
			}
		}
	}()

	return grp.Run()
}
