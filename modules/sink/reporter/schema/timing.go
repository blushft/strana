package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Timing holds the schema definition for the Timing entity.
type Timing struct {
	ent.Schema
}

// Fields of the Timing.
func (Timing) Fields() []ent.Field {
	return []ent.Field{
		field.String("category"),
		field.String("timing_label").StorageKey("label"),
		field.String("unit"),
		field.String("variable"),
		field.Float("value"),
	}
}

// Edges of the Timing.
func (Timing) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).Ref("timing"),
	}
}
