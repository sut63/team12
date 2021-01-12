package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// ClubappStatus holds the schema definition for the ClubappStatus entity.
type ClubappStatus struct {
	ent.Schema
}

// Fields of the ClubappStatus.
func (ClubappStatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("clubstatus").Unique(),
	}
}

// Edges of the ClubappStatus.
func (ClubappStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("clubapplication", Clubapplication.Type).StorageKey(edge.Column("clubappstatusID")),
	}
}
