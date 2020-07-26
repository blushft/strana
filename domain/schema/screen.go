package schema

import "github.com/facebookincubator/ent"

// Screen holds the schema definition for the Screen entity.
type Screen struct {
	ent.Schema
}

// Fields of the Screen.
func (Screen) Fields() []ent.Field {
	return nil
}

// Edges of the Screen.
func (Screen) Edges() []ent.Edge {
	return nil
}
