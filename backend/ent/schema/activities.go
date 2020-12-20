package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Activities holds the schema definition for the Activities entity.
type Activities struct {
	ent.Schema
}

// Fields of the Activities.
func (Activities) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("detail").NotEmpty(),
		field.Time("starttime"),
		field.Time("endtime"),
	}
}

// Edges of the Activities.
func (Activities) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("activitytype", ActivityType.Type).Ref("activities").Unique(),
		edge.From("academicyear", AcademicYear.Type).Ref("activities").Unique(),
	}
}



