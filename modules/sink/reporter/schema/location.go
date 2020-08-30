package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.String("street").Optional(),
		field.String("city").Optional(),
		field.String("state").Optional(),
		field.String("postalcode").Optional(),
		field.String("region").Optional(),
		field.String("locale").Optional(),
		field.String("country").Optional(),
		field.Float("longitude").Optional(),
		field.Float("latitude").Optional(),
		field.String("timezone").Optional(),
	}
}

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).Ref("location"),
	}
}
