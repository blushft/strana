package modules

import (
	"errors"
	"sync"

	"github.com/blushft/strana"
	"github.com/blushft/strana/platform/config"
)

var _globalRegistry = &registry{ctors: make(map[string]strana.ModuleConstructor)}

func Register(name string, mod strana.ModuleConstructor) {
	_globalRegistry.mu.Lock()
	_globalRegistry.ctors[name] = mod
	_globalRegistry.mu.Unlock()
}

type registry struct {
	mu    sync.Mutex
	ctors map[string]strana.ModuleConstructor
}

func (reg *registry) new(conf config.Module) (strana.Module, error) {
	reg.mu.Lock()
	defer reg.mu.Unlock()

	mod, ok := reg.ctors[conf.Type]
	if !ok {
		return nil, errors.New("no module found for " + conf.Type)
	}

	return mod(conf)
}

func New(conf config.Module) (strana.Module, error) {
	return _globalRegistry.new(conf)
}
