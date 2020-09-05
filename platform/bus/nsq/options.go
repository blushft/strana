package nsq

import (
	"github.com/blushft/strana/platform/logger"
	"github.com/nsqio/nsq/nsqd"
)

type Options struct {
	Logger interface{}
}

func nsqOptions(opts Options) *nsqd.Options {
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
