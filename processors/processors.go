package processors

import (
	"errors"
	"sync"

	"github.com/blushft/strana"
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
