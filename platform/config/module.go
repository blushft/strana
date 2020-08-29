package config

type Module struct {
	Name    string                 `json:"name" yaml:"name" mapstructure:"name"`
	Type    string                 `json:"type" yaml:"type" mapstructure:"type"`
	Source  MessagePath            `json:"source" yaml:"source" mapstructure:"source"`
	Sink    MessagePath            `json:"sink" yaml:"sink" mapstructure:"sink"`
	Options map[string]interface{} `json:"options" yaml:"options" mapstructure:"options"`
}

func DefaultModuleConfig() []Module {
	return []Module{
		{
			Name: "collector",
			Type: "collector",
			Sink: MessagePath{
				Topic: "collected_raw_message",
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
			Name: "loader",
			Type: "loader",
			Source: MessagePath{
				Topic: "collected_raw_message",
			},
		},
	}
}
