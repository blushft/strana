package contexts

import (
	"encoding/json"
	"net"

	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextNetwork event.ContextType = "network"

type Network struct {
	IP        net.IP `json:"ip,omitempty" structs:"ip,omitempty" mapstructure:"ip,omitempty"`
	UserAgent string `json:"userAgent,omitempty" structs:"userAgent,omitempty" mapstructure:"userAgent,omitempty"`
}

func NewNetwork(ip string) *Network {
	nip := net.ParseIP(ip)

	return &Network{
		IP: nip,
	}
}

func (ctx *Network) Type() event.ContextType {
	return ContextNetwork
}

func (ctx *Network) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *Network) Interface() interface{} {
	return ctx
}

func (ctx *Network) Validate() bool {
	if len(ctx.IP.String()) == 0 {
		return false
	}

	return true
}

func (ctx *Network) MarshalJSON() ([]byte, error) {
	ipt, err := ctx.IP.MarshalText()
	if err != nil {
		return nil, err
	}

	m := map[string]interface{}{
		"ip":        string(ipt),
		"userAgent": ctx.UserAgent,
	}

	return json.Marshal(m)
}

func (ctx *Network) UnmarshalJSON(b []byte) error {
	m := map[string]interface{}{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	ua, ok := m["userAgent"]
	if ok {
		ctx.UserAgent = ua.(string)
	}

	ipt, ok := m["ip"]
	if ok {
		ipb := ipt.(string)
		ip := net.IP{}
		if err := ip.UnmarshalText([]byte(ipb)); err != nil {
			return err
		}

		ctx.IP = ip
	}

	return nil
}
