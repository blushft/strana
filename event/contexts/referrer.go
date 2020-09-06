package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextReferrer event.ContextType = "referrer"

type Referrer struct {
	RefType  string `json:"type,omitempty" structs:"type,omitempty" mapstructure:"type,omitempty"`
	Name     string `json:"name,omitempty" structs:"name,omitempty" mapstructure:"name,omitempty"`
	Hostname string `json:"hostname,omitempty" structs:"hostname,omitempty" mapstructure:"hostname,omitempty"`
	Link     string `json:"link,omitempty" structs:"link,omitempty" mapstructure:"link,omitempty"`
}

func (ctx *Referrer) Type() event.ContextType {
	return ContextReferrer
}

func (ctx *Referrer) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *Referrer) Interface() interface{} {
	return ctx
}

func (ctx *Referrer) Validate() bool {
	if len(ctx.Hostname) == 0 {
		return false
	}

	return true
}
