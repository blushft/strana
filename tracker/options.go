package tracker

import "github.com/blushft/strana/pkg/event"

type Options struct {
	CollectorURL string
	AppInfo      *event.App
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
		event.WithContext(event.NewLibraryContext("go_tracker", "v0.0.1")),
	}

	if o.AppInfo != nil {
		evtOpts = append(evtOpts, event.WithAppContext(o.AppInfo))
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

func SetAppInfo(app *event.App) Option {
	return func(o *Options) {
		o.AppInfo = app
	}
}

func TrackingID(id string) Option {
	return func(o *Options) {
		o.TrackingID = id
	}
}
