package tracker

import "github.com/fatih/structs"

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
