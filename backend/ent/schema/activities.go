package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Activities holds the schema definition for the Activities entity.
type Activities struct {
	ent.Schema
}

// Fields of the Activities.
func (Activities) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("detail"),
		field.Time("starttime"),
		field.Time("endtime"),
	}
}

// Edges of the Activities.
func (Activities) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("activitytype", ActivityType.Type).Ref("activities").Unique(),
		edge.From("academicyear", AcademicYear.Type).Ref("activities").Unique(),
		edge.From("club",Club.Type).Ref("activities").Unique(),
	}
}



