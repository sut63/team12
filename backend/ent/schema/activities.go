package schema

import (
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Activities holds the schema definition for the Activities entity.
type Activities struct {
	ent.Schema
}

// Fields of the Activities.
func (Activities) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[A-Z][A-Za-z_]{6,50}$", s)
			if !match {
				return errors.New("รูปแบบชื่อกิจกรรมไม่ถูกต้อง")
			}
			return nil
		}),
		field.String("detail").Validate(func(s string) error {
			match, _ := regexp.MatchString(".{49,}", s)
			if !match {
				return errors.New("รูปแบบรายละเอียดกิจกรรมไม่ถูกต้อง")
			}
			return nil
		}),
		field.String("location").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[A-Z][A-Za-z_0-9]{2,49}$", s)
			if !match {
				return errors.New("รูปแบบชื่อสถานที่จัดกิจกรรมไม่ถูกต้อง")
			}
			return nil
		}),
		field.Time("starttime"),
		field.Time("endtime"),
	}
}

// Edges of the Activities.
func (Activities) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("activitytype", ActivityType.Type).Ref("activities").Unique(),
		edge.From("academicyear", AcademicYear.Type).Ref("activities").Unique(),
		edge.From("club",Club.Type).Ref("activities").Unique(),
	}
}



