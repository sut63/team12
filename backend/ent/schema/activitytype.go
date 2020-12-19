package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// ActivityType holds the schema definition for the ActivityType entity.
type ActivityType struct {
	ent.Schema
}

// Fields of the ActivityType.
func (ActivityType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the ActivityType.
func (ActivityType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("activities",Activities.Type).StorageKey(edge.Column("ActivityTypeID")),
	}
}
