package imports

import "reflect"

var Symbols = map[string]map[string]reflect.Value{}

//go:generate yaegi extract github.com/blushft/strana/event
//go:generate yaegi extract github.com/blushft/strana/event/contexts
//go:generate yaegi extract github.com/blushft/strana/event/events
//go:generate yaegi extract github.com/blushft/strana/processor
