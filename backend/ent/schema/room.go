package schema

import (
	"github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.String("room_name").NotEmpty(),
		field.String("room_location").NotEmpty(),
		field.String("room_floor").NotEmpty(),
		field.Int("max_contain").Positive(),
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roomuses", Roomuse.Type).StorageKey(edge.Column("room_id")),
	}
}
