package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextAlias event.ContextType = "alias"

type Alias struct {
	From string
	To   string
}

func NewAlias(from, to string) event.Context {
	return &Alias{
		From: from,
		To:   to,
	}
}

func (ctx *Alias) Type() event.ContextType {
	return ContextAlias
}

func (ctx *Alias) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *Alias) Interface() interface{} {
	return ctx
}

func (ctx *Alias) Validate() bool {
	if len(ctx.From) == 0 || len(ctx.To) == 0 {
		return false
	}

	return true
}
