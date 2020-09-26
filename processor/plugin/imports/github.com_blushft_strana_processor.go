// Code generated by 'github.com/containous/yaegi/extract github.com/blushft/strana/processor'. DO NOT EDIT.

// +build go1.14,!go1.15

package imports

import (
	"github.com/blushft/strana/event"
	"github.com/blushft/strana/processor"
	"reflect"
)

func init() {
	Symbols["github.com/blushft/strana/processor"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Execute": reflect.ValueOf(processor.Execute),

		// type definitions
		"Constructor":      reflect.ValueOf((*processor.Constructor)(nil)),
		"EventProcessor":   reflect.ValueOf((*processor.EventProcessor)(nil)),
		"ProcessFunc":      reflect.ValueOf((*processor.ProcessFunc)(nil)),
		"ProcessorWrapper": reflect.ValueOf((*processor.ProcessorWrapper)(nil)),

		// interface wrapper definitions
		"_EventProcessor": reflect.ValueOf((*_github_com_blushft_strana_processor_EventProcessor)(nil)),
	}
}

// _github_com_blushft_strana_processor_EventProcessor is an interface wrapper for EventProcessor type
type _github_com_blushft_strana_processor_EventProcessor struct {
	WProcess func(a0 *event.Event) ([]*event.Event, error)
}

func (W _github_com_blushft_strana_processor_EventProcessor) Process(a0 *event.Event) ([]*event.Event, error) {
	return W.WProcess(a0)
}