package useragent

import (
	"github.com/blushft/strana"
	"github.com/blushft/strana/pkg/event"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/processors"
	ua "github.com/mileusna/useragent"
)

func init() {
	processors.Register("useragent", func(config.Processor) (strana.Processor, error) {
		return &uaproc{
			validator: event.NewValidator(
				event.HasContext(event.ContextNetwork),
				event.ContextContains(event.ContextNetwork, "userAgent", true),
			),
		}, nil
	})
}

type uaproc struct {
	validator event.Validator
}

func (proc *uaproc) Process(evt *event.Event) ([]*event.Event, error) {
	if !proc.validator.Validate(evt) {
		return []*event.Event{evt}, nil
	}

	v := evt.Context[string(event.ContextNetwork)].Interface()
	netctx := v.(*event.Network)
	eua := ua.Parse(netctx.UserAgent)

	bctx := event.NewBrowserContext(&event.Browser{
		Name:      eua.Name,
		Version:   eua.Version,
		UserAgent: netctx.UserAgent,
	})

	osctx := event.NewOSContext(eua.OS, eua.OSVersion)

	devctx := event.NewDeviceContext(&event.Device{
		Mobile:  eua.Mobile,
		Tablet:  eua.Tablet,
		Desktop: eua.Desktop,
	})

	evt.SetContext(bctx)
	evt.SetContext(osctx)
	evt.SetContext(devctx)

	return []*event.Event{evt}, nil
}
