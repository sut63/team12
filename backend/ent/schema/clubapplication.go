package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Clubapplication holds the schema definition for the Clubapplication entity.
type Clubapplication struct {
	ent.Schema
}

// Fields of the Clubapplication.
func (Clubapplication) Fields() []ent.Field {
	return []ent.Field{
		field.String("contact").NotEmpty(),
		field.String("reason"),
		field.Time("addeddatetime"),
		field.String("addername").NotEmpty(),
		field.String("discipline").NotEmpty(),
		field.String("gender").NotEmpty(),
		field.Int("age").Positive(),
		field.Int("yaer").Positive(),
	}
}

// Edges of the Clubapplication.
func (Clubapplication) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("clubapplication").Unique(),
		edge.From("clubappstatus", ClubappStatus.Type).Ref("clubapplication").Unique(),
		edge.From("club", Club.Type).Ref("clubapplication").Unique(),
	}
}
