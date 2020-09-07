package processor

import (
	"log"

	"github.com/blushft/strana/event"
	"github.com/blushft/strana/platform/config"
)

type EventProcessor interface {
	Process(*event.Event) ([]*event.Event, error)
}

type ProcessFunc func(*event.Event) ([]*event.Event, error)
type ProcessorWrapper func(ProcessFunc) ProcessFunc

type Constructor func(conf config.Processor) (EventProcessor, error)

func Execute(procs []EventProcessor, evt *event.Event) ([]*event.Event, error) {
	q := []*event.Event{evt}

	for i := 0; len(q) > 0 && i < len(procs); i++ {
		var nextQ []*event.Event
		for _, m := range q {
			res, err := procs[i].Process(m)
			if err != nil {
				log.Printf("error executing processor: %s\n", err)
				return nil, err
			}

			nextQ = append(nextQ, res...)
		}

		q = nextQ
	}

	if len(q) == 0 {
		return nil, nil
	}

	return q, nil
}
