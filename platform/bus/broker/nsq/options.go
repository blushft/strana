package nsq

import (
	"github.com/blushft/strana/platform/bus/broker"
	"github.com/blushft/strana/platform/logger"
	"github.com/nsqio/nsq/nsqd"
)

type Options struct {
	BrokerOptions broker.Options
	NSQOptions    *nsqd.Options
}

func newOptions(opts broker.Options) Options {
	return Options{
		BrokerOptions: opts,
		NSQOptions:    nsqOptions(),
	}
}

func nsqOptions() *nsqd.Options {
	nopts := nsqd.NewOptions()

	nopts.Logger = nsqLogger{
		logger.Log(),
	}

	return nopts
}

type nsqLogger struct {
	log *logger.Logger
}

func (nl nsqLogger) Output(md int, msg string) error {
	nl.log.Info(msg)
	return nil
}
