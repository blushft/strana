package collector

import (
	"errors"

	"github.com/blushft/strana/platform"
	"github.com/blushft/strana/platform/config"
)

type Collector interface {
	platform.Module
	platform.Producer
}

func NewCollector(conf config.Collector) (Collector, error) {
	switch conf.Type {
	case "tracker":
		return newTrackingCollector(conf)
	default:
		return nil, errors.New("invalid collector type")
	}
}
