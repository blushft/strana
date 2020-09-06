package config

import "github.com/blushft/strana/platform/bus/message"

type Module struct {
	Name    string                 `json:"name" yaml:"name" mapstructure:"name"`
	Type    string                 `json:"type" yaml:"type" mapstructure:"type"`
	Source  message.Path           `json:"source" yaml:"source" mapstructure:"source"`
	Sink    message.Path           `json:"sink" yaml:"sink" mapstructure:"sink"`
	Options map[string]interface{} `json:"options" yaml:"options" mapstructure:"options"`
}

func DefaultModuleConfig() []Module {
	return []Module{
		{
			Name: "collector",
			Type: "collector",
			Sink: message.Path{
				Broker: "nsq",
				Topic:  "collected_raw_message",
			},
			Options: map[string]interface{}{
				"type": "tracker",
				"cache": map[string]interface{}{
					"default_expiration": 15,
				},
				"processors": []map[string]interface{}{
					{"name": "ua", "type": "useragent"},
				},
			},
		},
		{
			Name: "reporter",
			Type: "reporter",
			Source: message.Path{
				Broker:  "nsq",
				Channel: "reporter",
				Topic:   "collected_raw_message",
			},
		},
	}
}
