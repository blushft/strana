package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextOS event.ContextType = "os"

type OS struct {
	Name     string `json:"name" structs:"name" mapstructure:"name"`
	Family   string `json:"family,omitempty" structs:"family,omitempty" mapstructure:"family,omitempty"`
	Platform string `json:"platform,omitempty" structs:"platform,omitempty" mapstructure:"platform,omitempty"`
	Version  string `json:"version,omitempty" structs:"version,omitempty" mapstructure:"version,omitempty"`
}

func NewOSContext(name string) event.Context {
	return &OS{
		Name: name,
	}
}

func (ctx *OS) Type() event.ContextType {
	return ContextOS
}

func (ctx *OS) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *OS) Interface() interface{} {
	return ctx
}

func (ctx *OS) Validate() bool {
	if len(ctx.Name) == 0 {
		return false
	}

	return true
}
