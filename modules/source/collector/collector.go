package collector

import (
	"github.com/pkg/errors"

	"github.com/blushft/strana"
	"github.com/blushft/strana/modules"
	"github.com/blushft/strana/platform/config"
	"github.com/mitchellh/mapstructure"
)

func init() {
	modules.Register("collector", New)
}

type Collector interface {
	strana.Source
}

type Options struct {
	Type       string             `json:"type" yaml:"type" mapstructure:"type"`
	Cache      config.Cache       `json:"cache" yaml:"cache" mapstructure:"cache"`
	Processors []config.Processor `json:"processors" yaml:"processors" mapstructure:"processors"`
}

func New(conf config.Module) (strana.Module, error) {
	var opts Options
	if err := mapstructure.Decode(conf.Options, &opts); err != nil {
		return nil, errors.Wrap(err, "unable to decode collector options")
	}

	switch opts.Type {
	case "tracker":
		return newTrackingCollector(conf, opts)
	default:
		return nil, errors.New("invalid collector type")
	}
}
