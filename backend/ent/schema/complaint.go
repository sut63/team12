package schema

import (
	"github.com/facebook/ent"
)

// complaint holds the schema definition for the complaint entity.
type complaint struct {
	ent.Schema
}

// Fields of the complaint.
func (complaint) Fields() []ent.Field {
	return nil
}

// Edges of the complaint.
func (complaint) Edges() []ent.Edge {
	return nil
}
