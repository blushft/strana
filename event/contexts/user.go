package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextUser event.ContextType = "user"

type User struct {
	UserID string `json:"userID" structs:"userID" mapstructure:"userID"`
	AnonID string `json:"anonID" structs:"anonID" mapstructure:"anonID"`
	Name   string `json:"name,omitempty" structs:"name,omitempty" mapstructure:"name,omitempty"`
	Traits Traits `json:"traits,omitempty" structs:"traits,omitempty" mapstructure:"traits,omitempty"`
}

func (ctx *User) Type() event.ContextType {
	return ContextUser
}

func (ctx *User) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *User) Interface() interface{} {
	return ctx
}

func (ctx *User) Validate() bool {
	if len(ctx.UserID) == 0 || len(ctx.AnonID) == 0 {
		return false
	}

	return true
}

type Traits struct {
	Age         int                    `json:"age,omitempty" structs:"age,omitempty" mapstructure:"age,omitempty"`
	Birthday    string                 `json:"birthday,omitempty" structs:"birthday,omitempty" mapstructure:"birthday,omitempty"`
	Description string                 `json:"description,omitempty" structs:"description,omitempty" mapstructure:"description,omitempty"`
	Email       string                 `json:"email,omitempty" structs:"email,omitempty" mapstructure:"email,omitempty"`
	FirstName   string                 `json:"firstName,omitempty" structs:"firstName,omitempty" mapstructure:"firstName,omitempty"`
	LastName    string                 `json:"lastName,omitempty" structs:"lastName,omitempty" mapstructure:"lastName,omitempty"`
	Gender      string                 `json:"gender,omitempty" structs:"gender,omitempty" mapstructure:"gender,omitempty"`
	Phone       string                 `json:"phone,omitempty" structs:"phone,omitempty" mapstructure:"phone,omitempty"`
	Title       string                 `json:"title,omitempty" structs:"title,omitempty" mapstructure:"title,omitempty"`
	Username    string                 `json:"username,omitempty" structs:"username,omitempty" mapstructure:"username,omitempty"`
	Website     string                 `json:"website,omitempty" structs:"website,omitempty" mapstructure:"website,omitempty"`
	Extra       map[string]interface{} `json:"extra,omitempty" structs:"extra,omitempty" mapstructure:"extra,omitempty"`
}
