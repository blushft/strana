package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextTiming event.ContextType = "timing"

type Timing struct {
	Category string  `json:"category,omitempty" structs:"category,omitempty" mapstructure:"category,omitempty"`
	Label    string  `json:"label,omitempty" structs:"label,omitempty" mapstructure:"label,omitempty"`
	Unit     string  `json:"unit,omitempty" structs:"unit,omitempty" mapstructure:"unit,omitempty"`
	Variable string  `json:"variable,omitempty" structs:"variable,omitempty" mapstructure:"variable,omitempty"`
	Value    float64 `json:"value,omitempty" structs:"value,omitempty" mapstructure:"value,omitempty"`
}

func (ctx *Timing) Type() event.ContextType {
	return ContextTiming
}

func (ctx *Timing) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *Timing) Interface() interface{} {
	return ctx
}

func (ctx *Timing) Validate() bool {
	if ctx.Value == 0 {
		return false
	}

	return true
}
