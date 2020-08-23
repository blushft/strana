package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

// Fields of the App.
func (App) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("tracking_id").Unique().NotEmpty(),
	}
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sessions", Session.Type).Ref("app"),
		edge.From("pageviews", PageView.Type).Ref("app"),
		edge.To("stats", AppStat.Type),
		edge.To("page_stats", PageStat.Type),
	}
}
