package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Campaign holds the schema definition for the Campaign entity.
type Campaign struct {
	ent.Schema
}

// Fields of the Campaign.
func (Campaign) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("source").Optional(),
		field.String("medium").Optional(),
		field.String("term").Optional(),
		field.String("content").Optional(),
	}
}

// Edges of the Campaign.
func (Campaign) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("campaign"),
	}
}
