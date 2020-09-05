package useragent

import (
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/platform"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/processor"
	ua "github.com/mileusna/useragent"
)

func init() {
	platform.RegisterEventProcessor("useragent", func(config.Processor) (processor.EventProcessor, error) {
		return &uaproc{
			validator: event.NewValidator(
				event.WithRule("has_network", event.HasContext(event.ContextNetwork)),
				event.WithRule("has_user_agent", event.ContextContains(event.ContextNetwork, "userAgent", true)),
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
