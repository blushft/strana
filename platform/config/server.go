package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Server struct {
	Host          string       `json:"host" yaml:"host" mapstructure:"host" yaml.mapstructure:"host"`
	Port          string       `json:"port" yaml:"port" mapstructure:"port" yaml.mapstructure:"port"`
	EnableMetrics bool         `json:"enable_metrics" yaml:"enable_metrics" mapstructure:"enable_metrics" yaml.mapstructure:"enable_metrics"`
	Health        ServerHealth `json:"health" yaml:"health" mapstructure:"health"`
}

type ServerHealth struct {
	Enabled  bool   `json:"enabled" yaml:"enabled" mapstructure:"enabled"`
	Path     string `json:"path" yaml:"path" mapstructure:"path"`
	Listener string `json:"listener" yaml:"listener" mapstructure:"listener"`
}

func DefaultServerConfig() Server {
	v := viper.GetViper()
	return Server{
		Host: v.GetString("server.host"),
		Port: v.GetString("server.port"),
		Health: ServerHealth{
			Enabled: true,
			Path:    "/healthz",
		},
		EnableMetrics: false,
	}
}

func (c Server) HostPort() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
