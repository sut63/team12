package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// club holds the schema definition for the club entity.
type club struct {
	ent.Schema
}

// Fields of the club.
func (club) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").Positive(),
		field.String("name").NotEmpty(),
	}
}

// Edges of the club.
func (club) Edges() []ent.Edge {
	return nil
}

