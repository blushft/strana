package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextSession event.ContextType = "session"

type Session struct {
	ID string `json:"id" structs:"id" mapstructure:"id"`
}

func (ctx *Session) Type() event.ContextType {
	return ContextSession
}

func (ctx *Session) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *Session) Interface() interface{} {
	return ctx
}

func (ctx *Session) Validate() bool {
	if len(ctx.ID) == 0 {
		return false
	}

	return true
}
