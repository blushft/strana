package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextCampaign event.ContextType = "campaign"

type Campaign struct {
	Name    string `json:"name" structs:"name" mapstructure:"name"`
	Source  string `json:"source,omitempty" structs:"source,omitempty" mapstructure:"source,omitempty"`
	Medium  string `json:"medium,omitempty" structs:"medium,omitempty" mapstructure:"medium,omitempty"`
	Term    string `json:"term,omitempty" structs:"term,omitempty" mapstructure:"term,omitempty"`
	Content string `json:"content,omitempty" structs:"content,omitempty" mapstructure:"content,omitempty"`
}

func (ctx *Campaign) Type() event.ContextType {
	return ContextCampaign
}

func (ctx *Campaign) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *Campaign) Interface() interface{} {
	return ctx
}

func (ctx *Campaign) Validate() bool {
	if len(ctx.Name) == 0 {
		return false
	}

	return true
}
