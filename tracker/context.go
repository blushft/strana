package tracker

import (
	"time"

	"github.com/fatih/structs"
)

type Context interface {
	Values() map[string]interface{}
}

type Extra map[string]interface{}

func (e Extra) Values() map[string]interface{} {
	return e
}

type Action struct {
	Category string `json:"cat" mapstructure:"cat" structs:"cat"`
	Action   string `json:"act" mapstructure:"act" structs:"act"`
	Label    string `json:"lab" mapstructure:"lab" structs:"lab"`
	Property string `json:"prop" mapstructure:"prop" structs:"prop"`
	Value    string `json:"val" mapstructure:"val" structs:"val"`
}

type ActionOption func(*Action)

func NewAction(opts ...ActionOption) *Action {
	act := &Action{}

	for _, o := range opts {
		o(act)
	}

	return act
}

func (a *Action) Values() map[string]interface{} {
	return structs.Map(a)
}

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

type Device struct {
	OS         string `json:"os" mapstructure:"os"  structs:"os"`
	OSVersion  string `json:"osv" mapstructure:"osv"  structs:"osv"`
	Resolution string `json:"res" mapstructure:"res"  structs:"res"`
	Viewport   string `json:"vp" mapstructure:"vp"  structs:"vp"`
	ColorDepth string `json:"cd" mapstructure:"cd"  structs:"cd"`
	Timezone   string `json:"tz" mapstructure:"tz"  structs:"tz"`
	Language   string `json:"lang" mapstructure:"lang"  structs:"lang"`
	IPAddress  string `json:"ip" mapstructure:"ip"  structs:"ip"`
	UserAgent  string `json:"ua" mapstructure:"ua"  structs:"ua"`
	IsMobile   string `json:"mob" mapstructure:"mob"  structs:"mob"`
	IsDesktop  string `json:"dsk" mapstructure:"dsk"  structs:"dsk"`
	IsTablet   string `json:"tab" mapstructure:"tab"  structs:"tab"`
	IsBot      string `json:"bot" mapstructure:"bot"  structs:"bot"`
}

type DeviceOption func(*Device)

func NewDevice(opts ...DeviceOption) *Device {
	dev := &Device{}

	for _, o := range opts {
		o(dev)
	}

	return dev
}

func (d *Device) Values() map[string]interface{} {
	return structs.Map(d)
}

type Location struct {
	GeoID      string  `json:"geoid" mapstructure:"geoid" structs:"geoid"`
	Region     string  `json:"reg" mapstructure:"reg" structs:"reg"`
	Locale     string  `json:"loc" mapstructure:"loc" structs:"loc"`
	Country    string  `json:"cnty" mapstructure:"cnty" structs:"cnty"`
	City       string  `json:"city" mapstructure:"city" structs:"city"`
	PostalCode string  `json:"zip" mapstructure:"zip" structs:"zip"`
	Longitude  float64 `json:"long" mapstructure:"long" structs:"long"`
	Latitude   float64 `json:"lat" mapstructure:"lat" structs:"lat"`
}

type LocationOption func(*Location)

func NewLocation(opts ...LocationOption) *Location {
	l := &Location{}

	for _, o := range opts {
		o(l)
	}

	return l
}

func (l *Location) Values() map[string]interface{} {
	return structs.Map(l)
}

type Timing struct {
	Category string        `json:"utc" mapstructure:"utc" structs:"utc"`
	Variable string        `json:"utv" mapstructure:"utv" structs:"utv"`
	Label    string        `json:"utl" mapstructure:"utl" structs:"utl"`
	Time     time.Duration `json:"utt" mapstructure:"utt" structs:"utt"`
}

type TimingOption func(*Timing)

func NewTiming(opts ...TimingOption) *Timing {
	t := &Timing{}

	for _, o := range opts {
		o(t)
	}

	return t
}

func (t *Timing) Values() map[string]interface{} {
	return structs.Map(t)
}

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
