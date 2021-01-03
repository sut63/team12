package schema

import (
	"github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

// ClubType holds the schema definition for the ClubType entity.
type ClubType struct {
	ent.Schema
}

// Fields of the ClubType.
func (ClubType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the ClubType.
func (ClubType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("club", Club.Type).StorageKey(edge.Column("ClubType_ID")),
	}
}
