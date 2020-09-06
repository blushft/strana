package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

const ContextAction event.ContextType = "action"

type Action struct {
	Category string      `json:"category" structs:"category" mapstructure:"category"`
	Action   string      `json:"action" structs:"action" mapstructure:"action"`
	Label    string      `json:"label,omitempty" structs:"label,omitempty" mapstructure:"label,omitempty"`
	Property string      `json:"property,omitempty" structs:"property,omitempty" mapstructure:"property,omitempty"`
	Value    interface{} `json:"value,omitempty" structs:"value,omitempty" mapstructure:"value,omitempty"`
}

func newAction() event.Context {
	return &Action{}
}

func NewAction(cat, action string) event.Context {
	return &Action{
		Category: cat,
		Action:   action,
	}
}

func (ctx *Action) Type() event.ContextType {
	return ContextAction
}

func (ctx *Action) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *Action) Interface() interface{} {
	return ctx
}

func (ctx *Action) Validate() bool {
	if len(ctx.Category) == 0 || len(ctx.Action) == 0 {
		return false
	}

	return true
}

func decodeAction(v interface{}) (event.Context, error) {
	var ctx Action
	if err := mapstructure.Decode(v, &ctx); err != nil {
		return nil, err
	}

	return &ctx, nil
}
