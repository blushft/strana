package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextScreen event.ContextType = "screen"

type Screen struct {
	Name     string `json:"name" structs:"name" mapstructure:"name"`
	Category string `json:"category,omitempty" structs:"category,omitempty" mapstructure:"category,omitempty"`
}

func (ctx *Screen) Type() event.ContextType {
	return ContextScreen
}

func (ctx *Screen) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *Screen) Interface() interface{} {
	return ctx
}

func (ctx *Screen) Validate() bool {
	if len(ctx.Name) == 0 {
		return false
	}

	return true
}
