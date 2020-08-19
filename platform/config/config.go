package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Debug    bool     `json:"debug" yaml:"debug" mapstructure:"debug"`
	Logger   Logger   `json:"logger" yaml:"logger" mapstructure:"logger"`
	Database Database `json:"database" yaml:"database" mapstructure:"database"`
	Bus      Bus      `json:"bus" yaml:"bus" mapstructure:"bus"`
	Cache    Cache    `json:"cache" yaml:"cache" mapstructure:"cache"`
	Server   Server   `json:"server" yaml:"server" mapstructure:"server"`
	Modules  []Module `json:"modules" yaml:"modules" mapstructure:"modules"`
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
		Debug:    false,
		Logger:   DefaultLoggerConfig(),
		Database: DefaultDatabaseConfig(),
		Bus:      DefaultBusConfig(),
		Cache:    DefaultCacheConfig(),
		Server:   DefaultServerConfig(),
		Modules:  DefaultModuleConfig(),
	}
}
