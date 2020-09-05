package nats

import (
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

type Options struct {
	Port     int    `json:"port" structs:"port" mapstructure:"port"`
	HTTPPort int    `json:"httpPort" structs:"httpPort" mapstructure:"httpPort"`
	Token    string `json:"token" structs:"token" mapstructure:"token"`
}

func newOptions() Options {
	return Options{
		Port:     4442,
		HTTPPort: 4443,
		Token:    uuid.New().String(),
	}
}

func unmarshalOptions(m map[string]interface{}) (*Options, error) {
	options := newOptions()
	if err := mapstructure.Decode(m, &options); err != nil {
		return nil, err
	}

	return &options, nil
}
