// Code generated by 'github.com/containous/yaegi/extract github.com/blushft/strana/event'. DO NOT EDIT.

// +build go1.14,!go1.15

package imports

import (
	"github.com/blushft/strana/event"
	"reflect"
)

func init() {
	Symbols["github.com/blushft/strana/event"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Anonymous":       reflect.ValueOf(event.Anonymous),
		"Channel":         reflect.ValueOf(event.Channel),
		"ContextContains": reflect.ValueOf(event.ContextContains),
		"ContextInvalid":  reflect.ValueOf(event.ContextInvalid),
		"ContextValid":    reflect.ValueOf(event.ContextValid),
		"DeviceID":        reflect.ValueOf(event.DeviceID),
		"Empty":           reflect.ValueOf(event.Empty),
		"GetContextType":  reflect.ValueOf(event.GetContextType),
		"GroupID":         reflect.ValueOf(event.GroupID),
		"HasContext":      reflect.ValueOf(event.HasContext),
		"HasID":           reflect.ValueOf(event.HasID),
		"Interactive":     reflect.ValueOf(event.Interactive),
		"New":             reflect.ValueOf(event.New),
		"NewValidator":    reflect.ValueOf(event.NewValidator),
		"NonInteractive":  reflect.ValueOf(event.NonInteractive),
		"Platform":        reflect.ValueOf(event.Platform),
		"RegisterContext": reflect.ValueOf(event.RegisterContext),
		"RegisterType":    reflect.ValueOf(event.RegisterType),
		"SessionID":       reflect.ValueOf(event.SessionID),
		"TrackingID":      reflect.ValueOf(event.TrackingID),
		"UserID":          reflect.ValueOf(event.UserID),
		"WithContext":     reflect.ValueOf(event.WithContext),
		"WithContexts":    reflect.ValueOf(event.WithContexts),
		"WithRule":        reflect.ValueOf(event.WithRule),
		"WithRules":       reflect.ValueOf(event.WithRules),
		"WithValidator":   reflect.ValueOf(event.WithValidator),

		// type definitions
		"Context":           reflect.ValueOf((*event.Context)(nil)),
		"ContextContructor": reflect.ValueOf((*event.ContextContructor)(nil)),
		"ContextIterator":   reflect.ValueOf((*event.ContextIterator)(nil)),
		"ContextRegistry":   reflect.ValueOf((*event.ContextRegistry)(nil)),
		"ContextType":       reflect.ValueOf((*event.ContextType)(nil)),
		"Contexts":          reflect.ValueOf((*event.Contexts)(nil)),
		"Event":             reflect.ValueOf((*event.Event)(nil)),
		"EventConstructor":  reflect.ValueOf((*event.EventConstructor)(nil)),
		"EventRegistry":     reflect.ValueOf((*event.EventRegistry)(nil)),
		"Option":            reflect.ValueOf((*event.Option)(nil)),
		"Rule":              reflect.ValueOf((*event.Rule)(nil)),
		"Type":              reflect.ValueOf((*event.Type)(nil)),
		"Validator":         reflect.ValueOf((*event.Validator)(nil)),
		"ValidatorOption":   reflect.ValueOf((*event.ValidatorOption)(nil)),

		// interface wrapper definitions
		"_Context":         reflect.ValueOf((*_github_com_blushft_strana_event_Context)(nil)),
		"_ContextIterator": reflect.ValueOf((*_github_com_blushft_strana_event_ContextIterator)(nil)),
	}
}

// _github_com_blushft_strana_event_Context is an interface wrapper for Context type
type _github_com_blushft_strana_event_Context struct {
	WInterface func() interface{}
	WType      func() event.ContextType
	WValidate  func() bool
	WValues    func() map[string]interface{}
}

func (W _github_com_blushft_strana_event_Context) Interface() interface{}         { return W.WInterface() }
func (W _github_com_blushft_strana_event_Context) Type() event.ContextType        { return W.WType() }
func (W _github_com_blushft_strana_event_Context) Validate() bool                 { return W.WValidate() }
func (W _github_com_blushft_strana_event_Context) Values() map[string]interface{} { return W.WValues() }

// _github_com_blushft_strana_event_ContextIterator is an interface wrapper for ContextIterator type
type _github_com_blushft_strana_event_ContextIterator struct {
	WFirst func() event.Context
	WNext  func() event.Context
}

func (W _github_com_blushft_strana_event_ContextIterator) First() event.Context { return W.WFirst() }
func (W _github_com_blushft_strana_event_ContextIterator) Next() event.Context  { return W.WNext() }
