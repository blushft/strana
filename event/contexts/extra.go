package contexts

import (
	"github.com/blushft/strana/event"
)

const ContextExtra event.ContextType = "extra"

type Extra map[string]interface{}

func (ctx *Extra) Type() event.ContextType {
	return ContextExtra
}

func (ctx *Extra) Values() map[string]interface{} {
	return *ctx
}

func (ctx *Extra) Interface() interface{} {
	return ctx
}

func (ctx *Extra) Validate() bool {
	return true
}
