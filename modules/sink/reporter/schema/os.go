package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// OSContext holds the schema definition for the OSContext entity.
type OSContext struct {
	ent.Schema
}

func (OSContext) Config() ent.Config {
	return ent.Config{
		Table: "os",
	}
}

// Fields of the OSContext.
func (OSContext) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("family"),
		field.String("platform").Optional(),
		field.String("version"),
	}
}

// Edges of the OSContext.
func (OSContext) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).Ref("os"),
	}
}
