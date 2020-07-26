package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Debug     bool      `json:"debug" yaml:"debug" mapstructure:"debug"`
	Database  Database  `json:"database" yaml:"database" mapstructure:"database"`
	Bus       Bus       `json:"bus" yaml:"bus" mapstructure:"bus"`
	Cache     Cache     `json:"cache" yaml:"cache" mapstructure:"cache"`
	Server    Server    `json:"server" yaml:"server" mapstructure:"server"`
	Collector Collector `json:"collector" yaml:"collector" mapstructure:"collector"`
	Enhancer  Enhancer  `json:"enhancer" yaml:"enhancer" mapstructure:"enhancer"`
}

func NewConfig(v *viper.Viper) (*Config, error) {
	conf := DefaultConfig()
	if err := v.Unmarshal(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func DefaultConfig() Config {
	return Config{
		Debug:     false,
		Database:  DefaultDatabaseConfig(),
		Bus:       DefaultBusConfig(),
		Cache:     DefaultCacheConfig(),
		Server:    DefaultServerConfig(),
		Collector: DefaultCollectorConfig(),
		Enhancer:  DefaultEnhancerConfig(),
	}
}
