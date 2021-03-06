package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Action holds the schema definition for the Action entity.
type Action struct {
	ent.Schema
}

// Fields of the Action.
func (Action) Fields() []ent.Field {
	return []ent.Field{
		field.String("action"),
		field.String("category"),
		field.String("action_label").StorageKey("label").Optional(),
		field.String("property").Optional(),
		field.Bytes("value").Optional(),
	}
}

// Edges of the Action.
func (Action) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("action").Unique().Required(),
	}
}
