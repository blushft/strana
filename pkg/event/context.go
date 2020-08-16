package event

import (
	"encoding/json"
	"errors"

	"github.com/fatih/structs"
)

type ContextType string

const (
	ContextInvalid      ContextType = "invalid"
	ContextAction       ContextType = "action"
	ContextAlias        ContextType = "alias"
	ContextApp          ContextType = "app"
	ContextBrowser      ContextType = "browser"
	ContextCampaign     ContextType = "campaign"
	ContextConnectivity ContextType = "connectivity"
	ContextDevice       ContextType = "device"
	ContextExtra        ContextType = "extra"
	ContextLibrary      ContextType = "library"
	ContextLocation     ContextType = "location"
	ContextNetwork      ContextType = "network"
	ContextOS           ContextType = "os"
	ContextSession      ContextType = "session"
	ContextTiming       ContextType = "timing"
	ContextTraits       ContextType = "traits"
	ContextUser         ContextType = "user"
	ContextViewport     ContextType = "viewport"
)

func GetContextType(typ string) ContextType {
	switch typ {
	case "action":
		return ContextAction
	case "alias":
		return ContextAlias
	case "app":
		return ContextApp
	case "browser":
		return ContextBrowser
	case "campaign":
		return ContextCampaign
	case "connectivity":
		return ContextConnectivity
	case "device":
		return ContextDevice
	case "extra":
		return ContextExtra
	case "library":
		return ContextLibrary
	case "location":
		return ContextLocation
	case "network":
		return ContextNetwork
	case "os":
		return ContextOS
	case "session":
		return ContextSession
	case "timing":
		return ContextTiming
	case "traits":
		return ContextTraits
	case "user":
		return ContextUser
	case "viewport":
		return ContextViewport
	default:
		return ContextInvalid
	}
}

type Context interface {
	Type() ContextType
	Values() map[string]interface{}
	Interface() interface{}
}

type context struct {
	typ ContextType
	v   interface{}
}

func newContext(typ ContextType, v interface{}) Context {
	return &context{
		typ: typ,
		v:   v,
	}
}

func (ctx context) Type() ContextType {
	return ctx.typ
}

func (ctx context) Values() map[string]interface{} {
	return structs.Map(ctx.v)
}

func (ctx *context) Interface() interface{} {
	return ctx.v
}

type Contexts map[string]Context

func (c Contexts) Bind(v interface{}) {}

func (c Contexts) MarshalJSON() ([]byte, error) {
	cm := make(map[string]interface{}, len(c))
	for t, c := range c {
		cm[string(t)] = c.Values()
	}

	return json.Marshal(cm)
}

func (c *Contexts) UnmarshalJSON(b []byte) error {
	tc := make(Contexts)
	cm := make(map[string]json.RawMessage)

	if err := json.Unmarshal(b, &cm); err != nil {
		return err
	}

	for t, ce := range cm {
		nc, err := decodeContext(t, ce)
		if err != nil {
			return err
		}

		tc[t] = nc
	}

	*c = tc

	return nil
}

func emptyContext(typ ContextType) (Context, error) {
	if typ == ContextInvalid {
		return nil, errors.New("context type invalid")
	}

	switch typ {
	case ContextAction:
		return newContext(typ, &Action{}), nil
	case ContextAlias:
		return newContext(typ, &Alias{}), nil
	case ContextApp:
		return newContext(typ, &App{}), nil
	case ContextBrowser:
		return newContext(typ, &Browser{}), nil
	case ContextCampaign:
		return newContext(typ, &Campaign{}), nil
	case ContextConnectivity:
		return newContext(typ, &Connectivity{}), nil
	case ContextDevice:
		return newContext(typ, &Device{}), nil
	case ContextExtra:
		return newContext(typ, make(Extra)), nil
	case ContextLibrary:
		return newContext(typ, &Library{}), nil
	case ContextLocation:
		return newContext(typ, &Location{}), nil
	case ContextNetwork:
		return newContext(typ, &Network{}), nil
	case ContextOS:
		return newContext(typ, &OS{}), nil
	case ContextSession:
		return newContext(typ, &Session{}), nil
	case ContextTiming:
		return newContext(typ, &Timing{}), nil
	case ContextTraits:
		return newContext(typ, &Traits{}), nil
	case ContextUser:
		return newContext(typ, &User{}), nil
	case ContextViewport:
		return newContext(typ, &Viewport{}), nil
	}

	return nil, errors.New("context type unknown")
}

func decodeContext(typ string, vals json.RawMessage) (Context, error) {
	ct := GetContextType(typ)
	if ct == ContextInvalid {
		return nil, errors.New("context type invalid")
	}

	ec, err := emptyContext(ct)
	if err != nil {
		return nil, err
	}

	nv := ec.Interface()
	if err := json.Unmarshal(vals, &nv); err != nil {
		return nil, err
	}

	return newContext(ct, nv), nil
}
