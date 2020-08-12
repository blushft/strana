package tracker

import "github.com/fatih/structs"

type Web struct {
	URL      string `json:"url" mapstructure:"url" structs:"url"`
	Path     string `json:"path" mapstructure:"path" structs:"path"`
	Title    string `json:"page" mapstructure:"page" structs:"page"`
	Referrer string `json:"refr" mapstructure:"refr" structs:"refr"`
}

type WebOption func(*Web)

func NewWeb(opts ...WebOption) *Web {
	w := &Web{}

	for _, o := range opts {
		o(w)
	}

	return w
}

func (w *Web) Values() map[string]interface{} {
	return structs.Map(w)
}
