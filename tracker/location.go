package tracker

import "github.com/fatih/structs"

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
