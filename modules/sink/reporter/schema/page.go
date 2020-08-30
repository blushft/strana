package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Page holds the schema definition for the Page entity.
type Page struct {
	ent.Schema
}

// Fields of the Page.
func (Page) Fields() []ent.Field {
	return []ent.Field{
		field.String("hostname"),
		field.String("path"),
		field.String("referrer").Optional(),
		field.String("search").Optional(),
		field.String("title").Optional(),
		field.String("hash").Optional(),
	}
}

// Edges of the Page.
func (Page) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).Ref("page"),
	}
}
