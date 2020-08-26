package init

import (
	//import to register module constructors
	_ "github.com/blushft/strana/modules/broker/enhancer"
	_ "github.com/blushft/strana/modules/broker/sink"
	_ "github.com/blushft/strana/modules/sink/loader"
	_ "github.com/blushft/strana/modules/sink/reporter"
	_ "github.com/blushft/strana/modules/source/collector"
)
