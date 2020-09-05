package platform

import (
	"errors"
	"sync"

	"github.com/blushft/strana/processor"

	"github.com/blushft/strana/platform/bus"
	"github.com/blushft/strana/platform/bus/nsq"
	"github.com/blushft/strana/platform/config"
)

var ErrInvalidBroker = errors.New("invalid broker")

var (
	Brokers = map[string]func(...bus.Option) bus.Bus{
		"nsq": nsq.NewDefault,
	}
)

func NewBus(conf config.Bus) (bus.Bus, error) {
	switch conf.Broker {
	case "nsq":
		return nsq.New()
	default:
		return nil, ErrInvalidBroker
	}
}

var (
	_procRegistry = &procRegistry{reg: make(map[string]processor.Constructor)}
)

func RegisterEventProcessor(name string, proc processor.Constructor) {
	_procRegistry.mu.Lock()
	_procRegistry.reg[name] = proc
	_procRegistry.mu.Unlock()
}

func NewEventProcessor(conf config.Processor) (processor.EventProcessor, error) {
	return _procRegistry.new(conf)
}

func NewEventProcessorSet(conf []config.Processor) ([]processor.EventProcessor, error) {
	procs := make([]processor.EventProcessor, 0, len(conf))
	for _, p := range conf {
		proc, err := NewEventProcessor(p)
		if err != nil {
			return nil, err
		}

		procs = append(procs, proc)
	}

	return procs, nil
}

type procRegistry struct {
	mu  sync.Mutex
	reg map[string]processor.Constructor
}

func (reg *procRegistry) new(conf config.Processor) (processor.EventProcessor, error) {
	reg.mu.Lock()
	defer reg.mu.Unlock()

	p, ok := reg.reg[conf.Type]
	if !ok {
		return nil, errors.New("no processor found for " + conf.Type)
	}

	return p(conf)
}
