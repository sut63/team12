package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Usertype holds the schema definition for the Usertype entity.
type Usertype struct {
	ent.Schema
}

// Fields of the Usertype.
func (Usertype) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the Usertype.
func (Usertype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).StorageKey(edge.Column("UserTypeID")),
	}
}
