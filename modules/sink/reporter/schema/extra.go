package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Extra holds the schema definition for the Extra entity.
type Extra struct {
	ent.Schema
}

// Fields of the Extra.
func (Extra) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("values", map[string]interface{}{}),
	}
}

// Edges of the Extra.
func (Extra) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("extra"),
	}
}
