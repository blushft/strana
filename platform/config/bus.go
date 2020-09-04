package config

type Bus struct {
	Debug   bool                   `json:"debug" yaml:"debug" mapstructure:"debug" structs:"debug"`
	Trace   bool                   `json:"trace" yaml:"trace" mapstructure:"trace" structs:"trace"`
	Broker  string                 `json:"broker" structs:"broker" mapstructure:"broker"`
	Options map[string]interface{} `json:"options" structs:"options" mapstructure:"options"`
}

func DefaultBusConfig() Bus {
	return Bus{
		Debug:  false,
		Trace:  false,
		Broker: "nsq",
		Options: map[string]interface{}{
			"embedded": true,
		},
	}
}
