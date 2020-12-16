package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
)
 
// Useroom holds the schema definition for the User entity.
type Useroom struct {
   ent.Schema
}
 
// Fields of the Useroom.
func (Useroom) Fields() []ent.Field {
   return []ent.Field{
       field.String("name").NotEmpty(),
   }
}
 
// Edges of the Useroom.
func (Useroom) Edges() []ent.Edge {
   return nil
}

