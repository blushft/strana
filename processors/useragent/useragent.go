package useragent

import (
	"strconv"

	"github.com/blushft/strana"
	"github.com/blushft/strana/domain/entity"
	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/processors"
	ua "github.com/mileusna/useragent"
)

func init() {
	processors.Register("useragent", func(config.Processor) (strana.Processor, error) {
		return &uaproc{}, nil
	})
}

type uaproc struct{}

func (proc *uaproc) Process(msg *entity.RawMessage) ([]*entity.RawMessage, error) {
	v := ua.Parse(msg.UserAgent)

	msg.Browser = v.Name + " " + v.Version
	msg.BrowserName = v.Name
	msg.BrowserVersion = v.Version
	msg.OS = v.OS
	msg.OSVersion = v.OSVersion

	msg.IsMobile = strconv.FormatBool(v.Mobile)
	msg.IsDesktop = strconv.FormatBool(v.Desktop)
	msg.IsTablet = strconv.FormatBool(v.Tablet)
	msg.IsBot = strconv.FormatBool(v.Bot)

	return []*entity.RawMessage{msg}, nil
}
