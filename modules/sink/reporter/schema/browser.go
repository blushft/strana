package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Browser holds the schema definition for the Browser entity.
type Browser struct {
	ent.Schema
}

// Fields of the Browser.
func (Browser) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("version"),
		field.String("useragent").Optional(),
	}
}

// Edges of the Browser.
func (Browser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("browser").Unique(),
	}
}
