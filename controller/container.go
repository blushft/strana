package controller

import (
	"github.com/blushft/strana"
	"github.com/blushft/strana/modules"
	"github.com/blushft/strana/platform/bus"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/logger"
	"github.com/blushft/strana/platform/server"
	"github.com/blushft/strana/platform/store"
)

type container struct {
	svr   *server.Server
	bus   bus.Bus
	store *store.SQLStore
	log   *logger.Logger

	modc config.Module
	mod  strana.Module
}

func newContainers(conf config.Config, svr *server.Server, store *store.SQLStore, l *logger.Logger) (map[string]*container, error) {
	cs := make(map[string]*container, len(conf.Modules))

	for _, m := range conf.Modules {
		nc := &container{
			svr:   svr,
			store: store,
			log:   l,
			modc:  m,
		}

		if err := nc.initModule(); err != nil {
			return nil, err
		}

		cs[m.Name] = nc
	}

	return cs, nil
}

func (c *container) initModule() error {
	mod, err := modules.New(c.modc)
	if err != nil {
		return err
	}

	c.log.Mount(mod.Logger)

	if err := c.store.Mount(mod.Services); err != nil {
		return err
	}

	if err := c.svr.Mount(mod.Routes); err != nil {
		return err
	}

	c.mod = mod

	return nil
}

func (c *container) start(b bus.Bus) error {
	if err := b.Mount(c.mod); err != nil {
		return err
	}

	c.bus = b

	return nil
}

func (c *container) name() string {
	return c.modc.Name
}

func (c *container) module() strana.Module {
	return c.mod
}

func (c *container) t() string {
	return c.modc.Type
}
