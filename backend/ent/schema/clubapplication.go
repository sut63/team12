package schema

import (
	"time"

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

		field.String("applyname").NotEmpty(),

		field.String("contact"),
		field.String("reason"),
		field.Time("created_at").Default(time.Now),
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
