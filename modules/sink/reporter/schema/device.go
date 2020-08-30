package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("manufacturer").Optional(),
		field.String("model").Optional(),
		field.String("name").Optional(),
		field.String("type").Optional(),
		field.String("version").Optional(),
		field.Bool("mobile").Optional(),
		field.Bool("tablet").Optional(),
		field.Bool("desktop").Optional(),
		field.JSON("properties", map[string]interface{}{}).Optional(),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).Ref("device"),
	}
}
