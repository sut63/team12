package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/edge"
	"time"

)
// Borrow holds the schema definition for the Borrow entity.
type Borrow struct {
	ent.Schema
}

// Fields of the Borrow.
func (Borrow) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date").Default(time.Now).Immutable(),
	}
}

// Edges of the Borrow.
func (Borrow) Edges() []ent.Edge {
	return []ent.Edge{
		
	}
}




