package tracker

type Options struct {
	CollectorURL string
	AppID        int
	TrackingID   string
	QueueBuffer  int
}

func defaultOptions(opts ...Option) Options {
	options := Options{
		CollectorURL: "http://localhost:8863",
		AppID:        1,
		QueueBuffer:  25,
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

type Option func(*Options)

func (o Options) EventOptions() []EventOption {
	var evtOpts []EventOption

	if o.AppID > 0 {
		evtOpts = append(evtOpts, AppID(o.AppID))
	}

	if len(o.TrackingID) > 0 {
		evtOpts = append(evtOpts, TrackingID(o.TrackingID))
	}

	return evtOpts
}

func CollectorURL(u string) Option {
	return func(o *Options) {
		o.CollectorURL = u
	}
}

func WithAppID(id int) Option {
	return func(o *Options) {
		o.AppID = id
	}
}

func WithTrackingID(id string) Option {
	return func(o *Options) {
		o.TrackingID = id
	}
}
