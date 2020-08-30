package event

import (
	"encoding/json"
	"errors"
	"strings"

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
	ContextGroup        ContextType = "group"
	ContextLibrary      ContextType = "library"
	ContextLocation     ContextType = "location"
	ContextNetwork      ContextType = "network"
	ContextOS           ContextType = "os"
	ContextPage         ContextType = "page"
	ContextScreen       ContextType = "screen"
	ContextSession      ContextType = "session"
	ContextTiming       ContextType = "timing"
	ContextTraits       ContextType = "traits"
	ContextUser         ContextType = "user"
	ContextViewport     ContextType = "viewport"
)

func GetContextType(typ string) ContextType {
	for k := range registry {
		if strings.EqualFold(typ, string(k)) {
			return k
		}
	}

	return ContextInvalid
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

func (c Contexts) Map() map[string]interface{} {
	m := make(map[string]interface{}, len(c))
	for k, v := range c {
		m[k] = v
	}

	return m
}

func emptyContext(typ ContextType) (Context, error) {
	ctor, ok := registry[typ]
	if !ok {
		return nil, errors.New("context type unknown")
	}

	return ctor(), nil
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
