package filter

import "github.com/blushft/strana/event"

type Condition interface {
	Name() string
	Check(*event.Event) bool
}

type Conditions []Condition

func (cs Conditions) Add(cond Condition) {
	cs = append(cs, cond)
}
