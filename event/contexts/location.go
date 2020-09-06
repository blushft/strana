package contexts

import (
	"github.com/blushft/strana/event"
	"github.com/fatih/structs"
)

const ContextLocation event.ContextType = "location"

type Location struct {
	Street     string  `json:"street,omitempty" structs:"street,omitempty" mapstructure:"street,omitempty"`
	City       string  `json:"city,omitempty" structs:"city,omitempty" mapstructure:"city,omitempty"`
	State      string  `json:"state,omitempty" structs:"state,omitempty" mapstructure:"state,omitempty"`
	PostalCode string  `json:"postalCode,omitempty" structs:"postalCode,omitempty" mapstructure:"postalCode,omitempty"`
	Region     string  `json:"region,omitempty" structs:"region,omitempty" mapstructure:"region,omitempty"`
	Locale     string  `json:"locale,omitempty" structs:"locale,omitempty" mapstructure:"locale,omitempty"`
	Country    string  `json:"country,omitempty" structs:"country,omitempty" mapstructure:"country,omitempty"`
	Longitude  float64 `json:"longitude,omitempty" structs:"longitude,omitempty" mapstructure:"longitude,omitempty"`
	Latitude   float64 `json:"latitude,omitempty" structs:"latitude,omitempty" mapstructure:"latitude,omitempty"`
	Timezone   string  `json:"timezone,omitempty" structs:"timezone,omitempty" mapstructure:"timezone,omitempty"`
}

func (ctx *Location) Type() event.ContextType {
	return ContextLocation
}

func (ctx *Location) Values() map[string]interface{} {
	return structs.Map(ctx)
}

func (ctx *Location) Interface() interface{} {
	return ctx
}

func (ctx *Location) Validate() bool {
	return true
}
