package event

type Option func(*Event)

func TrackingID(id string) Option {
	return func(e *Event) {
		e.TrackingID = id
	}
}

func UserID(id string) Option {
	return func(e *Event) {
		e.UserID = id
	}
}

func GroupID(id string) Option {
	return func(e *Event) {
		e.GroupID = id
	}
}

func DeviceID(id string) Option {
	return func(e *Event) {
		e.DeviceID = id
	}
}

func SessionID(id string) Option {
	return func(e *Event) {
		e.SessionID = id
	}
}

func Anonymous(b bool) Option {
	return func(e *Event) {
		e.Anonymous = b
	}
}

func NonInteractive() Option {
	return func(e *Event) {
		e.NonInteractive = true
	}
}

func Interactive() Option {
	return func(e *Event) {
		e.NonInteractive = false
	}
}

func Platform(p string) Option {
	return func(e *Event) {
		e.Platform = p
	}
}

func Channel(c string) Option {
	return func(e *Event) {
		e.Channel = c
	}
}

func WithContext(ctx Context) Option {
	return func(e *Event) {
		e.SetContext(ctx)
	}
}

func WithContexts(ctx ...Context) Option {
	return func(e *Event) {
		e.SetContexts(ctx...)
	}
}

func WithValidator(v Validator) Option {
	return func(e *Event) {
		e.validator = v
	}
}
