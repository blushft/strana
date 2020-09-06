package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextConnectivity event.ContextType = "connectivity"

type Connectivity struct {
	Bluetooth bool   `json:"bluetooth,omitempty" structs:"bluetooth,omitempty" mapstructure:"bluetooth,omitempty"`
	Cellular  bool   `json:"cellular,omitempty" structs:"cellular,omitempty" mapstructure:"cellular,omitempty"`
	WIFI      bool   `json:"wifi,omitempty" structs:"wifi,omitempty" mapstructure:"wifi,omitempty"`
	Ethernet  bool   `json:"ethernet,omitempty" structs:"ethernet,omitempty" mapstructure:"ethernet,omitempty"`
	Carrier   string `json:"carrier,omitempty" structs:"carrier,omitempty" mapstructure:"carrier,omitempty"`
	ISP       string `json:"isp,omitempty" structs:"isp,omitempty" mapstructure:"isp,omitempty"`
}

func (ctx *Connectivity) Type() event.ContextType {
	return ContextConnectivity
}

func (ctx *Connectivity) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *Connectivity) Interface() interface{} {
	return ctx
}

func (ctx *Connectivity) Validate() bool {

	return true
}
