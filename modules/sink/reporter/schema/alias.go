package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Alias holds the schema definition for the Alias entity.
type Alias struct {
	ent.Schema
}

// Fields of the Alias.
func (Alias) Fields() []ent.Field {
	return []ent.Field{
		field.String("from"),
		field.String("to"),
	}
}

// Edges of the Alias.
func (Alias) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("alias").Unique(),
	}
}
