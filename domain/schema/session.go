package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.Bool("new_user"),
		field.Bool("is_unique"),
		field.Bool("is_bounce"),
		field.Bool("is_finished"),
		field.Int("duration").Optional(),
		field.Time("started_at"),
		field.Time("finished_at").Nillable().Optional(),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app", App.Type).Unique().Required(),
		edge.To("user", User.Type).Unique().Required(),
		edge.To("device", Device.Type).Unique(),
		edge.From("pageviews", PageView.Type).Ref("session"),
	}
}
