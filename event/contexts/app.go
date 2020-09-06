package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextApp event.ContextType = "app"

type App struct {
	Name       string                 `json:"name" structs:"name" mapstructure:"name"`
	Version    string                 `json:"version,omitempty" structs:"version,omitempty" mapstructure:"version,omitempty"`
	Build      string                 `json:"build,omitempty" structs:"build,omitempty" mapstructure:"build,omitempty"`
	Namespace  string                 `json:"namespace,omitempty" structs:"namespace,omitempty" mapstructure:"namespace,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty" structs:"properties,omitempty" mapstructure:"properties,omitempty"`
}

func NewApp(n string) event.Context {
	return &App{
		Name: n,
	}
}

func (ctx *App) Type() event.ContextType {
	return ContextApp
}

func (ctx *App) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *App) Interface() interface{} {
	return ctx
}

func (ctx *App) Validate() bool {
	if len(ctx.Name) == 0 {
		return false
	}

	return true
}
