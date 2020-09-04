package controller

import (
	"github.com/blushft/strana/platform/bus"
	"github.com/blushft/strana/platform/bus/nsq"
)

var (
	Brokers = map[string]func(...bus.Option) bus.Bus{
		"nsq": nsq.NewDefault,
	}
)
