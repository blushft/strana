package init

import (
	// Registration side effects
	_ "github.com/blushft/strana/processors/geoip"
	_ "github.com/blushft/strana/processors/log"
	_ "github.com/blushft/strana/processors/useragent"
)
