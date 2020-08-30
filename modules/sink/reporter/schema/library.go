package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Library holds the schema definition for the Library entity.
type Library struct {
	ent.Schema
}

// Fields of the Library.
func (Library) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("version").Optional(),
	}
}

// Edges of the Library.
func (Library) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).Ref("library"),
	}
}
