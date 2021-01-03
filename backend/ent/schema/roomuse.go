package schema

import (
	"github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

// Roomuse holds the schema definition for the Roomuse entity.
type Roomuse struct {
	ent.Schema
}

// Fields of the Roomuse.
func (Roomuse) Fields() []ent.Field {
	return []ent.Field{
		field.Time("added_time"),
	}
}

// Edges of the Roomuse.
func (Roomuse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rooms", Room.Type).Ref("roomuses").Unique(),
		edge.From("purposes", Roompurpose.Type).Ref("roompurpose").Unique(),
		edge.From("users", User.Type).Ref("Roomuse").Unique(),

	}
}
