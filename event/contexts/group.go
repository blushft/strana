package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextGroup event.ContextType = "group"

type Group struct {
	ID   string `json:"id,omitempty" structs:"id,omitempty" mapstructure:"id,omitempty"`
	Name string `json:"name,omitempty" structs:"name,omitempty" mapstructure:"name,omitempty"`
}

func (ctx *Group) Type() event.ContextType {
	return ContextGroup
}

func (ctx *Group) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *Group) Interface() interface{} {
	return ctx
}

func (ctx *Group) Validate() bool {
	if len(ctx.ID) == 0 || len(ctx.Name) == 0 {
		return false
	}

	return true
}
