package schema

import (
 issue-27
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/edge"
	"time"

)

	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

 main
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
 issue-27
	return []ent.Edge{
		
	}
}


	return []ent.Edge{}
} main
