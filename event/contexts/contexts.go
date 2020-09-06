package contexts

import "github.com/blushft/strana/event"

func init() {
	registerContexts()
}

func registerContexts() {
	event.RegisterContext(ctxCtor(ContextAction))
	event.RegisterContext(ctxCtor(ContextAlias))
	event.RegisterContext(ctxCtor(ContextApp))
	event.RegisterContext(ctxCtor(ContextBrowser))
	event.RegisterContext(ctxCtor(ContextCampaign))
	event.RegisterContext(ctxCtor(ContextConnectivity))
	event.RegisterContext(ctxCtor(ContextDevice))
	event.RegisterContext(ctxCtor(ContextExtra))
	event.RegisterContext(ctxCtor(ContextDevice))
	event.RegisterContext(ctxCtor(ContextExtra))
	event.RegisterContext(ctxCtor(ContextGroup))
	event.RegisterContext(ctxCtor(ContextLibrary))
	event.RegisterContext(ctxCtor(ContextLocation))
	event.RegisterContext(ctxCtor(ContextNetwork))
	event.RegisterContext(ctxCtor(ContextOS))
	event.RegisterContext(ctxCtor(ContextPage))
	event.RegisterContext(ctxCtor(ContextReferrer))
	event.RegisterContext(ctxCtor(ContextScreen))
	event.RegisterContext(ctxCtor(ContextSession))
	event.RegisterContext(ctxCtor(ContextTiming))
	event.RegisterContext(ctxCtor(ContextUser))
	event.RegisterContext(ctxCtor(ContextViewport))
}

func ctxCtor(typ event.ContextType) (event.ContextType, func() event.Context) {
	return typ, func() event.Context { return newContext(typ) }
}

func newContext(typ event.ContextType) event.Context {
	switch typ {
	case ContextAction:
		return &Action{}
	case ContextAlias:
		return &Alias{}
	case ContextApp:
		return &App{}
	case ContextBrowser:
		return &Browser{}
	case ContextCampaign:
		return &Campaign{}
	case ContextConnectivity:
		return &Connectivity{}
	case ContextDevice:
		return &Device{}
	case ContextExtra:
		return &Extra{}
	case ContextGroup:
		return &Group{}
	case ContextLibrary:
		return &Library{}
	case ContextLocation:
		return &Location{}
	case ContextNetwork:
		return &Network{}
	case ContextOS:
		return &OS{}
	case ContextPage:
		return &Page{}
	case ContextSession:
		return &Session{}
	case ContextTiming:
		return &Timing{}
	case ContextUser:
		return &User{}
	case ContextViewport:
		return &Viewport{}
	default:
		return nil
	}
}
