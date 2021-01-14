package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// UserStatus holds the schema definition for the UserStatus entity.
type UserStatus struct {
	ent.Schema
}

// Fields of the UserStatus.
func (UserStatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("userstatus").Unique(),
	}
}

// Edges of the UserStatus.
func (UserStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).StorageKey(edge.Column("userstatus_id")),
	}
}
