package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.Int("age").Positive(),
		field.String("email").NotEmpty(),
		field.String("password"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("usertype", Usertype.Type).Ref("user").Unique(),
		edge.From("FromClub", Club.Type).Ref("userclub").Unique(),

		edge.From("gender", Gender.Type).Ref("users").Unique(),
		edge.From("userstatus", UserStatus.Type).Ref("users").Unique(),
		edge.From("discipline", Discipline.Type).Ref("users").Unique(),
		edge.From("year", Year.Type).Ref("users").Unique(),

		edge.To("club", Club.Type).StorageKey(edge.Column("UserID")),
		edge.To("clubapplication", Clubapplication.Type).StorageKey(edge.Column("UserID")),
		edge.To("UserToComplaint", Complaint.Type).StorageKey(edge.Column("UserID")),
		edge.To("Roomuse", Roomuse.Type).StorageKey(edge.Column("UserID")),
	}
}
