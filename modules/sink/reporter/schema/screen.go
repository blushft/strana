package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Screen holds the schema definition for the Screen entity.
type Screen struct {
	ent.Schema
}

// Fields of the Screen.
func (Screen) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("category").Optional(),
	}
}

// Edges of the Screen.
func (Screen) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).Ref("screen"),
	}
}
