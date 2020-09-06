package tracker

import (
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/event/contexts"
)

type Options struct {
	CollectorURL string
	AppInfo      *contexts.App
	Platform     string
	TrackingID   string
	QueueBuffer  int
}

func defaultOptions(opts ...Option) Options {
	options := Options{
		CollectorURL: "http://localhost:8863",
		QueueBuffer:  25,
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

type Option func(*Options)

func (o Options) EventOptions() []event.Option {
	evtOpts := []event.Option{
		event.WithContext(&contexts.Library{Name: "go_tracker", Version: "v0.0.1"}),
	}

	if o.AppInfo != nil {
		evtOpts = append(evtOpts, event.WithContext(o.AppInfo))
	}

	if len(o.TrackingID) > 0 {
		evtOpts = append(evtOpts, event.TrackingID(o.TrackingID))
	}

	if len(o.Platform) > 0 {
		evtOpts = append(evtOpts, event.Platform(o.Platform))
	} else {
		evtOpts = append(evtOpts, event.Platform("srv"))
	}

	return evtOpts
}

func CollectorURL(u string) Option {
	return func(o *Options) {
		o.CollectorURL = u
	}
}

func SetAppInfo(app *contexts.App) Option {
	return func(o *Options) {
		o.AppInfo = app
	}
}

func TrackingID(id string) Option {
	return func(o *Options) {
		o.TrackingID = id
	}
}
