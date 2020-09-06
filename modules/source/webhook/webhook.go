package webhook

import (
	"github.com/blushft/strana"
	"github.com/blushft/strana/modules"
	"github.com/blushft/strana/platform/config"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

func init() {
	modules.Register("webhook", New)
}

type Webhook interface {
	strana.Source
}

type Options struct {
	Hooks      map[string]HookOptions `json:"hooks" structs:"hooks" mapstructure:"hooks"`
	Processors []config.Processor     `json:"processors" structs:"processors" mapstructure:"processors"`
}

type HookOptions map[string]interface{}

func New(conf config.Module) (strana.Module, error) {
	var opts Options
	if err := mapstructure.Decode(conf.Options, &opts); err != nil {
		return nil, errors.Wrap(err, "decoding webhook options")
	}

	for t, o := range opts.Hooks {
		switch t {
		case "pingdom":
			return newPingdomHook(conf, o, opts.Processors...)
		default:
			return nil, errors.New("invalid webhook type")
		}
	}

	return nil, errors.New("invalid webhook configuration")
}
