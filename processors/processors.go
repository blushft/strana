package processors

import (
	"errors"
	"log"
	"sync"

	"github.com/blushft/strana"
	"github.com/blushft/strana/pkg/event"
	"github.com/blushft/strana/platform/config"
)

var _globalRegistry = &registry{reg: make(map[string]strana.ProcessorConstructor)}

func Register(name string, proc strana.ProcessorConstructor) {
	_globalRegistry.mu.Lock()
	_globalRegistry.reg[name] = proc
	_globalRegistry.mu.Unlock()
}

func New(conf config.Processor) (strana.Processor, error) {
	return _globalRegistry.new(conf)
}

func Execute(procs []strana.Processor, evt *event.Event) ([]*event.Event, error) {
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

type registry struct {
	mu  sync.Mutex
	reg map[string]strana.ProcessorConstructor
}

func (reg *registry) new(conf config.Processor) (strana.Processor, error) {
	reg.mu.Lock()
	defer reg.mu.Unlock()

	p, ok := reg.reg[conf.Type]
	if !ok {
		return nil, errors.New("no processor found for " + conf.Type)
	}

	return p(conf)
}
