package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextPage event.ContextType = "page"

type Page struct {
	Hash     string `json:"hash,omitempty" structs:"hash,omitempty" mapstructure:"hash,omitempty"`
	Path     string `json:"path,omitempty" structs:"path,omitempty" mapstructure:"path,omitempty"`
	Referrer string `json:"referrer,omitempty" structs:"referrer,omitempty" mapstructure:"referrer,omitempty"`
	Search   string `json:"search,omitempty" structs:"search,omitempty" mapstructure:"search,omitempty"`
	Title    string `json:"title,omitempty" structs:"title,omitempty" mapstructure:"title,omitempty"`
	Hostname string `json:"hostname" structs:"hostname" mapstructure:"hostname"`
}

func (ctx *Page) Type() event.ContextType {
	return ContextPage
}

func (ctx *Page) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *Page) Interface() interface{} {
	return ctx
}

func (ctx *Page) Validate() bool {
	if len(ctx.Hostname) == 0 {
		return false
	}

	return true
}
