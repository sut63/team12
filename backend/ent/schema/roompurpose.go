package schema

import (
	"github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

// Roompurpose holds the schema definition for the Roompurpose entity.
type Roompurpose struct {
	ent.Schema
}

// Fields of the Roompurpose.
func (Roompurpose) Fields() []ent.Field {
	return []ent.Field{
		field.String("purpose").NotEmpty(),
	}
}

// Edges of the Roompurpose.
func (Roompurpose) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roompurpose", Roomuse.Type).StorageKey(edge.Column("purpose_id")),
	}
}
