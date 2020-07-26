package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// PageStat holds the schema definition for the PageStat entity.
type PageStat struct {
	ent.Schema
}

// Fields of the PageStat.
func (PageStat) Fields() []ent.Field {
	return []ent.Field{
		field.Int("pageviews"),
		field.Int("visitors"),
		field.Int("entries"),
		field.Float("bounce_rate"),
		field.Int("known_durations").Default(0),
		field.Float("avg_duration"),
		field.Time("date"),
	}
}

// Edges of the PageStat.
func (PageStat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app", App.Type).Ref("page_stats").Required().Unique(),
		edge.From("hostname", Hostname.Type).Ref("page_stats").Required().Unique(),
		edge.From("pathname", Pathname.Type).Ref("page_stats").Required().Unique(),
	}
}
