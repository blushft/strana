package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"github.com/google/uuid"
)

// RawEvent holds the schema definition for the RawEvent entity.
type RawEvent struct {
	ent.Schema
}

// Fields of the RawEvent.
func (RawEvent) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("tracking_id", uuid.UUID{}),
		field.String("user_id"),
		field.Bool("anonymous"),
		field.String("group_id").Optional(),
		field.String("session_id").Optional(),
		field.String("device_id").Optional(),
		field.String("event"),
		field.Bool("non_interactive"),
		field.String("channel").Optional(),
		field.String("platform").Optional(),
		field.Time("timestamp"),
		field.JSON("context", map[string]interface{}{}),
	}
}

// Edges of the RawEvent.
func (RawEvent) Edges() []ent.Edge {
	return nil
}

func (RawEvent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "tracking_id", "group_id", "session_id", "device_id", "event"),
	}
}
