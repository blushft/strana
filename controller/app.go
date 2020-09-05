package controller

import (
	"github.com/blushft/strana/platform"
	"github.com/blushft/strana/platform/bus"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
	"github.com/blushft/strana/platform/server"
	"github.com/blushft/strana/platform/store"
	"github.com/oklog/run"
)

type App struct {
	conf  config.Config
	svr   *server.Server
	bus   bus.Bus
	store *store.SQLStore
	log   *logger.Logger

	modules map[string]*container
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

	mods, err := newContainers(conf, svr, s, l)
	if err != nil {
		return nil, err
	}

	a := &App{
		conf:    conf,
		svr:     svr,
		bus:     bus,
		store:   s,
		log:     l,
		modules: mods,
	}

	svr.Mount(a.routes)

	return a, nil
}

func (a *App) Start() error {
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
		for k, c := range a.modules {
			a.log.Infof("mounting events for module %s", k)

			if err := a.bus.Mount(c.module()); err != nil {
				logger.Log().Fatal(err.Error())
			}
		}
	}()

	return grp.Run()
}
