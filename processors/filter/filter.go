package filter

import "github.com/blushft/strana/event"

type filter struct {
	conds Conditions
}

func (f *filter) Process(evt *event.Event) ([]*event.Event, error) {
	for _, c := range f.conds {
		if c.Check(evt) {
			return nil, nil
		}
	}

	return []*event.Event{evt}, nil
}
