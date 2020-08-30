package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Referrer holds the schema definition for the Referrer entity.
type Referrer struct {
	ent.Schema
}

// Fields of the Referrer.
func (Referrer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("type").Optional(),
		field.String("hostname").Optional(),
		field.String("link").Optional(),
	}
}

// Edges of the Referrer.
func (Referrer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).Ref("referrer"),
	}
}
