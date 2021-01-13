package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Purpose holds the schema definition for the Purpose entity.
type Purpose struct {
	ent.Schema
}

// Fields of the Purpose.
func (Purpose) Fields() []ent.Field {
	return []ent.Field{
		field.String("purpose").NotEmpty(),
	}
}

// Edges of the Purpose.
func (Purpose) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roomuses", Roomuse.Type).StorageKey(edge.Column("purpose_id")),
	}
}

