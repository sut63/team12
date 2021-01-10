package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Club holds the schema definition for the Club entity.
type Club struct {
	ent.Schema
}

// Fields of the Club.
func (Club) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("purpose").NotEmpty(),
	}
}

// Edges of the Club.
func (Club) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("club").Unique(),
		edge.From("clubtype", ClubType.Type).Ref("club").Unique(),
		edge.From("clubbranch", ClubBranch.Type).Ref("club").Unique(),
		
		edge.To("clubapplication", Clubapplication.Type).StorageKey(edge.Column("ClubID")),
		edge.To("ClubToComplaint", Complaint.Type).StorageKey(edge.Column("ClubID")),
		edge.To("activities",Activities.Type).StorageKey(edge.Column("ClubID")),
	 	edge.To("userclub",User.Type).StorageKey(edge.Column("ClubID")), 
	}
}
