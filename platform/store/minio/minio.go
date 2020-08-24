package minio

type Options struct {
	Bucket string `json:"bucket" yaml:"bucket"`
	Prefix string `json:"prefix" yaml:"prefix"`
	URL    string `json:"url" yaml:"url"`
	Access string `json:"access" yaml:"access"`
	Secret string `json:"secret" yaml:"secret"`
	Secure bool   `json:"secure" yaml:"secure"`
}

type Option func(*Options)

type Store struct {
	opts Options
}
