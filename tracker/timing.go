package tracker

import (
	"time"

	"github.com/fatih/structs"
)

type Timing struct {
	Category string        `json:"utc" mapstructure:"utc" structs:"utc"`
	Variable string        `json:"utv" mapstructure:"utv" structs:"utv"`
	Label    string        `json:"utl" mapstructure:"utl" structs:"utl"`
	Time     time.Duration `json:"utt" mapstructure:"utt" structs:"utt"`
}

type TimingOption func(*Timing)

func NewTiming(opts ...TimingOption) *Timing {
	t := &Timing{}

	for _, o := range opts {
		o(t)
	}

	return t
}

func (t *Timing) Values() map[string]interface{} {
	return structs.Map(t)
}
