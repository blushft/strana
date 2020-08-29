package config

import (
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

const (
	DefaultBusPort     = 4442
	DefaultBusHTTPPort = 4443
)

type Bus struct {
	Debug    bool              `json:"debug" yaml:"debug" mapstructure:"debug" structs:"debug"`
	Trace    bool              `json:"trace" yaml:"trace" mapstructure:"trace" structs:"trace"`
	Port     int               `json:"port" structs:"port" mapstructure:"port"`
	HTTPPort int               `json:"httpPort" structs:"httpPort" mapstructure:"httpPort"`
	Token    string            `json:"token" structs:"token" mapstructure:"token"`
	Brokers  map[string]PubSub `json:"brokers" structs:"brokers" mapstructure:"brokers"`
}

func DefaultBusConfig() Bus {
	return Bus{
		Debug:    false,
		Trace:    false,
		Port:     DefaultBusPort,
		HTTPPort: DefaultBusHTTPPort,
		Token:    uuid.New().String(),
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
