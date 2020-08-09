package modules

import (
	"errors"

	"github.com/blushft/strana"
	"github.com/blushft/strana/modules/collector"
	"github.com/blushft/strana/modules/enhancer"
	"github.com/blushft/strana/modules/loader"
	"github.com/blushft/strana/platform/config"
)

func New(conf config.Module) (strana.Module, error) {
	switch conf.Type {
	case "collector":
		return collector.New(conf)
	case "enhancer":
		return enhancer.New(conf)
	case "loader":
		return loader.New(conf)
	default:
		return nil, errors.New("invalid module: " + conf.Type)
	}
}
