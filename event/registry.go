package event

var ctxreg ContextRegistry

type ContextContructor func() Context
type ContextRegistry map[ContextType]ContextContructor

func init() {
	ctxreg = make(map[ContextType]ContextContructor)
}

func contextCtor(typ ContextType, v interface{}) ContextContructor {
	return func() Context {
		return newContext(typ, v)
	}
}

func RegisterContext(typ ContextType, ctor ContextContructor) {
	ctxreg[typ] = ctor
}
