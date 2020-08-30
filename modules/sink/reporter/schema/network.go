package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Network holds the schema definition for the Network entity.
type Network struct {
	ent.Schema
}

// Fields of the Network.
func (Network) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip"),
		field.String("useragent").Optional(),
	}
}

// Edges of the Network.
func (Network) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).Ref("network"),
	}
}
