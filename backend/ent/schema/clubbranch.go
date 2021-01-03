package schema

import (
	"github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

// ClubBranch holds the schema definition for the ClubBranch entity.
type ClubBranch struct {
	ent.Schema
}

// Fields of the ClubBranch.
func (ClubBranch) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the ClubBranch.
func (ClubBranch) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("club", Club.Type).StorageKey(edge.Column("ClubBranch_ID")),
	}
}
