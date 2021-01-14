package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Complaint holds the schema definition for the Complaint entity.
type Complaint struct {
	ent.Schema
}

// Fields of the Complaint.
func (Complaint) Fields() []ent.Field {
	return []ent.Field{
		field.String("info"),
		field.String("date"),
	}
}

// Edges of the Complaint.
func (Complaint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ComplaintToUser", User.Type).
			Ref("UserToComplaint").
			Unique(),
		edge.From("ComplaintToComplaintType", ComplaintType.Type).
			Ref("ComplaintTypeToComplaint").
			Unique(),
		edge.From("ComplaintToClub", Club.Type).
			Ref("ClubToComplaint").
			Unique(),
	}
}
