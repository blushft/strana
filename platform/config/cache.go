package config

type Cache struct {
	Store             string `json:"store" yaml:"store" mapstructure:"store"`
	DefaultExpiration int    `json:"default_expiration" yaml:"default_expiration" mapstructure:"default_expiration"`
}

func DefaultCacheConfig() Cache {
	return Cache{
		Store:             "memory",
		DefaultExpiration: 15,
	}
}
