package bus

import "github.com/blushft/strana"

type Bus interface {
	strana.EventHandler
	Mount(strana.Module) error
	Start() error
	Shutdown() error
}
