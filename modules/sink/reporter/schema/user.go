package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.Bool("is_anonymous"),
		field.String("name").Optional(),
		field.String("title").Optional(),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.String("email").Optional(),
		field.String("username").Optional(),
		field.Int("age").Optional(),
		field.Time("birthday").Optional().Nillable(),
		field.Enum("gender").NamedValues("Male", "M", "Female", "F", "Other", "O").Optional(),
		field.String("phone").Optional(),
		field.String("website").Optional(),
		field.JSON("extra", map[string]interface{}{}).Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("aliases", Alias.Type),
		edge.From("events", Event.Type).Ref("user"),
		edge.From("groups", Group.Type).Ref("users"),
	}
}
