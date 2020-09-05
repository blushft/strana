package config

type Bus struct {
	Debug   bool                   `json:"debug" yaml:"debug" mapstructure:"debug" structs:"debug"`
	Trace   bool                   `json:"trace" yaml:"trace" mapstructure:"trace" structs:"trace"`
	Brokers []Broker               `json:"broker" structs:"broker" mapstructure:"broker"`
	Options map[string]interface{} `json:"options" structs:"options" mapstructure:"options"`
}

func DefaultBusConfig() Bus {
	return Bus{
		Debug: false,
		Trace: false,
		Brokers: []Broker{
			{
				Name: "nsq",
				Type: "nsq",
				Options: map[string]interface{}{
					"embedded": true,
				},
			},
		},
	}
}

type Broker struct {
	Name    string                 `json:"name" structs:"name" mapstructure:"name"`
	Type    string                 `json:"type" structs:"type" mapstructure:"type"`
	Options map[string]interface{} `json:"options" structs:"options" mapstructure:"options"`
}
