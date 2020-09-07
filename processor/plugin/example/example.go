package example

import (
	"fmt"

	"github.com/blushft/strana/event"
	"github.com/blushft/strana/processor"
)

type Config struct {
	Name string
}

type SimplePlugin struct {
	config *Config
	count  int
}

func NewConfig() *Config {
	return &Config{}
}

func New(c *Config) processor.EventProcessor {
	return &SimplePlugin{config: c}
}

func (p *SimplePlugin) Process(evt *event.Event) ([]*event.Event, error) {
	p.count++
	fmt.Printf("plugin %s, event count: %d\n", p.config.Name, p.count)

	var res []*event.Event

	res = append(res, evt)

	return res, nil
}
