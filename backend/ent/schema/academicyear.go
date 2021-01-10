package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// AcademicYear holds the schema definition for the AcademicYear entity.
type AcademicYear struct {
	ent.Schema
}

// Fields of the AcademicYear.
func (AcademicYear) Fields() []ent.Field {
	return []ent.Field{
		field.String("semester").NotEmpty(),
	}
}

// Edges of the AcademicYear.
func (AcademicYear) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("activities",Activities.Type).StorageKey(edge.Column("AcademicYearID")),
	}
}
