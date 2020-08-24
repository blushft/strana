package config

type FileStore struct {
	Provider string                 `json:"provider" yaml:"provider"`
	Options  map[string]interface{} `json:"options" yaml:"options"`
}
