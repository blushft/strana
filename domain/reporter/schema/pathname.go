package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Pathname holds the schema definition for the Pathname entity.
type Pathname struct {
	ent.Schema
}

// Fields of the Pathname.
func (Pathname) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the Pathname.
func (Pathname) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("page_stats", PageStat.Type),
	}
}
