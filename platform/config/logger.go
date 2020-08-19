package config

type Logger struct {
	Outputs map[string]Output `json:"outputs" yaml:"outputs" mapstructure:"outputs"`
}

type Output struct {
	Type    string                 `json:"type" yaml:"type" mapstructure:"type"`
	Options map[string]interface{} `json:"options" yaml:"options" mapstructure:"options"`
}

func DefaultLoggerConfig() Logger {
	return Logger{
		Outputs: map[string]Output{
			"console": {Type: "text"},
		},
	}
}
