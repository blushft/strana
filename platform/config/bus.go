package config

import "github.com/mitchellh/mapstructure"

type Bus struct {
	Debug   bool `json:"debug" yaml:"debug" mapstructure:"debug"`
	Trace   bool `json:"trace" yaml:"trace" mapstructure:"trace"`
	Brokers map[string]PubSub
}

func DefaultBusConfig() Bus {
	return Bus{
		Debug: false,
		Trace: false,
		Brokers: map[string]PubSub{
			"in_process": DefaultPubSubConfig(),
		},
	}
}

type PubSub struct {
	Driver string                 `json:"driver" yaml:"driver" mapstructure:"driver"`
	Config map[string]interface{} `json:"config" yaml:"config" mapstructure:"config"`
}

func DefaultPubSubConfig() PubSub {
	return PubSub{
		Driver: "memory",
		Config: map[string]interface{}{
			"OutputChannelBuffer":           0,
			"BlockPublishUntilSubcriberAck": false,
		},
	}
}

func (ps PubSub) DriverConfig(v interface{}) error {
	return mapstructure.Decode(ps.Config, v)
}

type MessagePath struct {
	Source string
	Topic  string
	Broker string
}
