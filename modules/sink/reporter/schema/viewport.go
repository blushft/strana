package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Viewport holds the schema definition for the Viewport entity.
type Viewport struct {
	ent.Schema
}

// Fields of the Viewport.
func (Viewport) Fields() []ent.Field {
	return []ent.Field{
		field.Int("density"),
		field.Int("width"),
		field.Int("height"),
	}
}

// Edges of the Viewport.
func (Viewport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).Ref("viewport"),
	}
}
