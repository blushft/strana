package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

// Fields of the App.
func (App) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("version").Optional(),
		field.String("build").Optional(),
		field.String("namespace").Optional(),
		field.JSON("properties", map[string]interface{}{}).Optional(),
	}
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).Ref("app"),
	}
}

func (App) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "version", "build").Unique(),
	}
}
