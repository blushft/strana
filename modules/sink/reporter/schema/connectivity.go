package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Connectivity holds the schema definition for the Connectivity entity.
type Connectivity struct {
	ent.Schema
}

// Fields of the Connectivity.
func (Connectivity) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("bluetooth"),
		field.Bool("cellular"),
		field.Bool("wifi"),
		field.Bool("ethernet"),
		field.Bool("carrier"),
		field.Bool("isp"),
	}
}

// Edges of the Connectivity.
func (Connectivity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("connectivity").Unique(),
	}
}
