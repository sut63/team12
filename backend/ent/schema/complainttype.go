package schema

import (
	"github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

// ComplaintType holds the schema definition for the ComplaintType entity.
type ComplaintType struct {
	ent.Schema
}

// Fields of the ComplaintType.
func (ComplaintType) Fields() []ent.Field {
	return []ent.Field{
		field.String("description"),
	}
}

// Edges of the ComplaintType.
func (ComplaintType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ComplaintTypeToComplaint", Complaint.Type).StorageKey(edge.Column("Type")),
	}
}
