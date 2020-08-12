package tracker

import "github.com/fatih/structs"

type Device struct {
	OS         string `json:"os"`
	OSVersion  string `json:"osv"`
	Resolution string `json:"res"`
	Viewport   string `json:"vp"`
	ColorDepth string `json:"cd"`
	Timezone   string `json:"tz"`
	Language   string `json:"lang"`
	IPAddress  string `json:"ip"`
	UserAgent  string `json:"ua"`
	IsMobile   string `json:"mob"`
	IsDesktop  string `json:"dsk"`
	IsTablet   string `json:"tab"`
	IsBot      string `json:"bot"`
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
