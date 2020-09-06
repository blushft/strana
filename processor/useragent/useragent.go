package useragent

import (
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/event/contexts"
	"github.com/blushft/strana/platform"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/processor"
	ua "github.com/mileusna/useragent"
)

func init() {
	platform.RegisterEventProcessor("useragent", func(config.Processor) (processor.EventProcessor, error) {
		return &uaproc{
			validator: event.NewValidator(
				event.WithRule("has_network", event.HasContext(contexts.ContextNetwork)),
				event.WithRule("has_user_agent", event.ContextContains(contexts.ContextNetwork, "userAgent", true)),
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

	v := evt.Context[string(contexts.ContextNetwork)].Interface()
	netctx := v.(*contexts.Network)
	eua := ua.Parse(netctx.UserAgent)

	bctx := &contexts.Browser{
		Name:      eua.Name,
		Version:   eua.Version,
		UserAgent: netctx.UserAgent,
	}

	osctx := &contexts.OS{
		Name:    eua.OS,
		Version: eua.OSVersion,
	}

	devctx := &contexts.Device{
		Mobile:  eua.Mobile,
		Tablet:  eua.Tablet,
		Desktop: eua.Desktop,
	}

	evt.SetContext(bctx)
	evt.SetContext(osctx)
	evt.SetContext(devctx)

	return []*event.Event{evt}, nil
}
