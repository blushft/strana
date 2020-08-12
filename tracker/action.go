package tracker

import "github.com/fatih/structs"

type Action struct {
	Category string `json:"cat" mapstructure:"cat"`
	Action   string `json:"act" mapstructure:"act"`
	Label    string `json:"lab" mapstructure:"lab"`
	Property string ` json:"prop" mapstructure:"prop"`
	Value    string `json:"val" mapstructure:"val"`
}

type ActionOption func(*Action)

func NewAction(opts ...ActionOption) *Action {
	act := &Action{}

	for _, o := range opts {
		o(act)
	}

	return act
}

func (a *Action) Values() map[string]interface{} {
	return structs.Map(a)
}
