package config

import "github.com/mitchellh/mapstructure"

func BindOptions(m map[string]interface{}, v interface{}) error {
	cfg := &mapstructure.DecoderConfig{
		DecodeHook:       mapstructure.StringToSliceHookFunc(","),
		WeaklyTypedInput: true,
		Result:           v,
	}

	dec, err := mapstructure.NewDecoder(cfg)
	if err != nil {
		return err
	}

	return dec.Decode(m)
}
