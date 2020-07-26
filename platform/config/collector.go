package config

type Collector struct {
	Type      string `json:"type" yaml:"type" mapstructure:"type"`
	Publisher Source `json:"sink" yaml:"sink" mapstructure:"sink"`
	Cache     Cache  `json:"cache" yaml:"cache" mapstructure:"cache"`
}

func DefaultCollectorConfig() Collector {
	return Collector{
		Type: "tracker",
		Publisher: Source{
			Module: "collector",
			Topic:  "collected_raw_message",
			Broker: "in_process",
		},
		Cache: DefaultCacheConfig(),
	}
}
