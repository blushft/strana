package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("tracking_id"),
		field.Enum("event").
			Values(
				"action",
				"alias",
				"group",
				"identify",
				"pageview",
				"screenview",
				"session",
				"timing",
				"transaction",
			),
		field.Bool("non_interactive"),
		field.String("channel").Optional(),
		field.String("platform").Optional(),
		field.JSON("properties", map[string]interface{}{}).Optional(),
		field.Time("timestamp"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("action", Action.Type).Unique(),
		edge.To("alias", Alias.Type).Unique(),
		edge.To("app", App.Type).Unique(),
		edge.To("browser", Browser.Type).Unique(),
		edge.To("campaign", Campaign.Type).Unique(),
		edge.To("connectivity", Connectivity.Type).Unique(),
		edge.To("device", Device.Type).Unique(),
		edge.To("extra", Extra.Type).Unique(),
		edge.To("group", Group.Type).Unique(),
		edge.To("library", Library.Type).Unique(),
		edge.To("location", Location.Type).Unique(),
		edge.To("network", Network.Type).Unique(),
		edge.To("os", OSContext.Type).Unique(),
		edge.To("page", Page.Type).Unique(),
		edge.To("referrer", Referrer.Type).Unique(),
		edge.To("screen", Screen.Type).Unique(),
		edge.To("session", Session.Type).Unique(),
		edge.To("timing", Timing.Type).Unique(),
		edge.To("viewport", Viewport.Type).Unique(),
		edge.To("user", User.Type).Unique(),
	}
}
