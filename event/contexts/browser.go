package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextBrowser event.ContextType = "browser"

type Browser struct {
	Name      string `json:"name,omitempty" structs:"name,omitempty" mapstructure:"name,omitempty"`
	Version   string `json:"version,omitempty" structs:"version,omitempty" mapstructure:"version,omitempty"`
	UserAgent string `json:"userAgent,omitempty" structs:"userAgent,omitempty" mapstructure:"userAgent,omitempty"`
}

func NewBrowser(n string) event.Context {
	return &Browser{
		Name: n,
	}
}

func (ctx *Browser) Type() event.ContextType {
	return ContextBrowser
}

func (ctx *Browser) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *Browser) Interface() interface{} {
	return ctx
}

func (ctx *Browser) Validate() bool {
	if len(ctx.Name) == 0 {
		return false
	}

	return true
}
