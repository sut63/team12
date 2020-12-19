package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
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
