package tracker

import "github.com/fatih/structs"

type Browser struct {
	Browser        string `json:"b" mapstructure:"b" structs:"b"`
	BrowserName    string `json:"bn" mapstructure:"bn" structs:"bn"`
	BrowserVersion string `json:"bv" mapstructure:"bv" structs:"bv"`
	BrowserEngine  string `json:"be" mapstructure:"be" structs:"be"`
}

type BrowserOption func(*Browser)

func NewBrowser(opts ...BrowserOption) *Browser {
	br := &Browser{}

	for _, o := range opts {
		o(br)
	}

	return br
}

func (b *Browser) Values() map[string]interface{} {
	return structs.Map(b)
}
