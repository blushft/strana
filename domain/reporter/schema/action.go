package schema

import "github.com/facebook/ent"

// Action holds the schema definition for the Action entity.
type Action struct {
	ent.Schema
}

// Fields of the Action.
func (Action) Fields() []ent.Field {
	return nil
}

// Edges of the Action.
func (Action) Edges() []ent.Edge {
	return nil
}
