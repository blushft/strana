package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextDevice event.ContextType = "device"

type Device struct {
	ID           string                 `json:"id,omitempty" structs:"id,omitempty" mapstructure:"id,omitempty"`
	Manufacturer string                 `json:"manufacturer,omitempty" structs:"manufacturer,omitempty" mapstructure:"manufacturer,omitempty"`
	Model        string                 `json:"model,omitempty" structs:"model,omitempty" mapstructure:"model,omitempty"`
	Name         string                 `json:"name,omitempty" structs:"name,omitempty" mapstructure:"name,omitempty"`
	Kind         string                 `json:"type,omitempty" structs:"type,omitempty" mapstructure:"type,omitempty"`
	Version      string                 `json:"version,omitempty" structs:"version,omitempty" mapstructure:"version,omitempty"`
	Mobile       bool                   `json:"mobile,omitempty" structs:"mobile,omitempty" mapstructure:"mobile,omitempty"`
	Tablet       bool                   `json:"tablet,omitempty" structs:"tablet,omitempty" mapstructure:"tablet,omitempty"`
	Desktop      bool                   `json:"desktop,omitempty" structs:"desktop,omitempty" mapstructure:"desktop,omitempty"`
	Properties   map[string]interface{} `json:"properties,omitempty" structs:"properties,omitempty" mapstructure:"properties,omitempty"`
}

func NewDeviceContext(id string) event.Context {
	return &Device{
		ID: id,
	}
}

func (ctx *Device) Type() event.ContextType {
	return ContextDevice
}

func (ctx *Device) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *Device) Interface() interface{} {
	return ctx
}

func (ctx *Device) Validate() bool {
	return true
}
