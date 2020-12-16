package schema

import "github.com/facebookincubator/ent"

// Clubapplication holds the schema definition for the Clubapplication entity.
type Clubapplication struct {
	ent.Schema
}

// Fields of the Clubapplication.
func (Clubapplication) Fields() []ent.Field {
	return nil
}

// Edges of the Clubapplication.
func (Clubapplication) Edges() []ent.Edge {
	return nil
}
