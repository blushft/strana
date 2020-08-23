package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// PageView holds the schema definition for the PageView entity.
type PageView struct {
	ent.Schema
}

// Fields of the PageView.
func (PageView) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("hostname"),
		field.String("pathname"),
		field.String("referrer"),
		field.Bool("is_entry"),
		field.Bool("is_finished"),
		field.Int("duration"),
		field.Time("timestamp"),
		field.String("user_agent").Optional(),
		field.String("ip_address").Optional(),
		field.String("screen_dim").Optional(),
		field.JSON("extra", map[string]interface{}{}).Optional(),
	}
}

// Edges of the PageView.
func (PageView) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app", App.Type).Unique(),
		edge.To("session", Session.Type).Unique(),
		edge.To("user", User.Type).Unique(),
	}
}
