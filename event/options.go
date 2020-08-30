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
		e.Context[string(ctx.Type())] = ctx
	}
}

func WithActionContext(a *Action) Option {
	return WithContext(newContext(ContextAction, a))
}

func WithAppContext(a *App) Option {
	return WithContext(newContext(ContextApp, a))
}

func WithAliasContext(a *Alias) Option {
	return WithContext(newContext(ContextAlias, a))
}

func WithBrowserContext(b *Browser) Option {
	return WithContext(newContext(ContextBrowser, b))
}

func WithCampaignContext(c *Campaign) Option {
	return WithContext(newContext(ContextCampaign, c))
}

func WithConnectivityContext(c *Connectivity) Option {
	return WithContext(newContext(ContextConnectivity, c))
}

func WithDeviceContext(d *Device) Option {
	return WithContext(newContext(ContextDevice, d))
}

func WithExtraContext(e *Extra) Option {
	return WithContext(newContext(ContextExtra, e))
}

func WithGroupContext(g *Group) Option {
	return WithContext(newContext(ContextGroup, g))
}

func WithLocationContext(l *Location) Option {
	return WithContext(newContext(ContextLocation, l))
}

func WithNetworkContext(n *Network) Option {
	return WithContext(newContext(ContextNetwork, n))
}

func WithOSContext(o *OS) Option {
	return WithContext(newContext(ContextOS, o))
}

func WithPageContext(p *Page) Option {
	return WithContext(newContext(ContextPage, p))
}

func WithSessionContext(s *Session) Option {
	return WithContext(newContext(ContextSession, s))
}

func WithTimingContext(t *Timing) Option {
	return WithContext(newContext(ContextTiming, t))
}

func WithUserContext(u *User) Option {
	return WithContext(NewUserContext(u))
}

func WithViewportContext(v *Viewport) Option {
	return WithContext(newContext(ContextViewport, v))
}
