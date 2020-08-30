package event

var registry ContextRegistry

type ContextContructor func() Context
type ContextRegistry map[ContextType]ContextContructor

func init() {
	registry = make(map[ContextType]ContextContructor)
	registerContexts()
}

func contextCtor(typ ContextType, v interface{}) ContextContructor {
	return func() Context {
		return newContext(typ, v)
	}
}

func registerContexts() {
	registry[ContextAction] = contextCtor(ContextAction, &Action{})
	registry[ContextAlias] = contextCtor(ContextAlias, &Alias{})
	registry[ContextApp] = contextCtor(ContextApp, &App{})
	registry[ContextBrowser] = contextCtor(ContextBrowser, &Browser{})
	registry[ContextCampaign] = contextCtor(ContextCampaign, &Campaign{})
	registry[ContextConnectivity] = contextCtor(ContextConnectivity, &Connectivity{})
	registry[ContextDevice] = contextCtor(ContextDevice, &Device{})
	registry[ContextExtra] = contextCtor(ContextExtra, make(Extra))
	registry[ContextGroup] = contextCtor(ContextGroup, &Group{})
	registry[ContextLibrary] = contextCtor(ContextLibrary, &Library{})
	registry[ContextLocation] = contextCtor(ContextLocation, &Location{})
	registry[ContextNetwork] = contextCtor(ContextNetwork, &Network{})
	registry[ContextOS] = contextCtor(ContextOS, &OS{})
	registry[ContextSession] = contextCtor(ContextSession, &Session{})
	registry[ContextTiming] = contextCtor(ContextTiming, &Timing{})
	registry[ContextTraits] = contextCtor(ContextTraits, &Traits{})
	registry[ContextUser] = contextCtor(ContextUser, &User{})
	registry[ContextViewport] = contextCtor(ContextViewport, &Viewport{})
}

func RegisterContext(typ ContextType, ctor ContextContructor) {
	registry[typ] = ctor
}
