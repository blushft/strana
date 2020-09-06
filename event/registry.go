package event

var evtreg EventRegistry
var ctxreg ContextRegistry

func init() {
	ctxreg = make(map[ContextType]ContextContructor)
	evtreg = make(map[Type]bool)
}

type EventConstructor func(...Option) *Event
type EventRegistry map[Type]bool

func RegisterType(typ Type) {
	evtreg[typ] = true
}

type ContextContructor func() Context
type ContextRegistry map[ContextType]ContextContructor

func RegisterContext(typ ContextType, ctor ContextContructor) {
	ctxreg[typ] = ctor
}
