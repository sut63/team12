package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Discipline holds the schema definition for the Discipline entity.
type Discipline struct {
	ent.Schema
}

// Fields of the Discipline.
func (Discipline) Fields() []ent.Field {
	return []ent.Field{
		field.String("discipline").Unique(),
	}
}

// Edges of the Discipline.
func (Discipline) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).StorageKey(edge.Column("discipline_id")),
	}
}
