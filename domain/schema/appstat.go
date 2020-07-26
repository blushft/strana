package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// AppStat holds the schema definition for the AppStat entity.
type AppStat struct {
	ent.Schema
}

// Fields of the AppStat.
func (AppStat) Fields() []ent.Field {
	return []ent.Field{
		field.Int("pageviews"),
		field.Int("visitors"),
		field.Int("sessions"),
		field.Float("bouce_rate"),
		field.Int("known_durations"),
		field.Float("avg_duration"),
		field.Time("date"),
	}
}

// Edges of the AppStat.
func (AppStat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app", App.Type).Ref("stats").Unique(),
	}
}
