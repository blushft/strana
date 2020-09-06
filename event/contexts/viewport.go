package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextViewport event.ContextType = "viewport"

type Viewport struct {
	Density int `json:"density,omitempty" structs:"density,omitempty" mapstructure:"density,omitempty"`
	Width   int `json:"width,omitempty" structs:"width,omitempty" mapstructure:"width,omitempty"`
	Height  int `json:"height,omitempty" structs:"height,omitempty" mapstructure:"height,omitempty"`
}

func (ctx *Viewport) Type() event.ContextType {
	return ContextViewport
}

func (ctx *Viewport) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *Viewport) Interface() interface{} {
	return ctx
}

func (ctx *Viewport) Validate() bool {

	return true
}
