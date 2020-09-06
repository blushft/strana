package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextLibrary event.ContextType = "library"

type Library struct {
	Name    string `json:"name" structs:"name" mapstructure:"name"`
	Version string `json:"version,omitempty" structs:"version,omitempty" mapstructure:"version,omitempty"`
}

func (ctx *Library) Type() event.ContextType {
	return ContextLibrary
}

func (ctx *Library) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *Library) Interface() interface{} {
	return ctx
}

func (ctx *Library) Validate() bool {
	return true
}
