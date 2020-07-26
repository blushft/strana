package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Hostname holds the schema definition for the Hostname entity.
type Hostname struct {
	ent.Schema
}

// Fields of the Hostname.
func (Hostname) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the Hostname.
func (Hostname) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("page_stats", PageStat.Type),
	}
}
