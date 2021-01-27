package schema

import (
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Clubapplication holds the schema definition for the Clubapplication entity.
type Clubapplication struct {
	ent.Schema
}

// Fields of the Clubapplication.
func (Clubapplication) Fields() []ent.Field {
	return []ent.Field{
		field.String("contact").Validate(func(s string) error {
			match, _ := regexp.MatchString(".{10,}", s)
			if !match {
				return errors.New("ข้อมูลติดต่อผิดพลาด กรุณาป้อนข้อมูลอยา่งน้อย 10 ตัวอักษร")
			}
			return nil
		}),
		field.String("reason"),
		field.Time("addeddatetime"),
		field.String("addername").Validate(func(n string) error {
			match, _ := regexp.MatchString("^[A-Za-zก-๙]+[ \t\r\n\v\f]+[A-Za-zก-๙]+[^๐-๙]$", n)
			if !match {
				return errors.New("รูปแบบรายชื่อผิดพลาด กรุณาป้อน ชื่อ วรรค นามสกุล อย่างน้อย 1 ตัว")
			}
			return nil
		}),
		field.String("discipline").NotEmpty(),
		field.String("gender").NotEmpty(),
		field.Int("age").Min(1).Max(200),
		field.Int("yaer").Positive(),
	}
}

// Edges of the Clubapplication.
func (Clubapplication) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("clubapplication").Unique(),
		edge.From("clubappstatus", ClubappStatus.Type).Ref("clubapplication").Unique(),
		edge.From("club", Club.Type).Ref("clubapplication").Unique(),
	}
}
