// Code generated by entc, DO NOT EDIT.

package academicyear

import (
	"github.com/OMENX/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Semester applies equality check predicate on the "semester" field. It's identical to SemesterEQ.
func Semester(v string) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemester), v))
	})
}

// SemesterEQ applies the EQ predicate on the "semester" field.
func SemesterEQ(v string) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemester), v))
	})
}

// SemesterNEQ applies the NEQ predicate on the "semester" field.
func SemesterNEQ(v string) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSemester), v))
	})
}

// SemesterIn applies the In predicate on the "semester" field.
func SemesterIn(vs ...string) predicate.AcademicYear {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AcademicYear(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSemester), v...))
	})
}

// SemesterNotIn applies the NotIn predicate on the "semester" field.
func SemesterNotIn(vs ...string) predicate.AcademicYear {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AcademicYear(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSemester), v...))
	})
}

// SemesterGT applies the GT predicate on the "semester" field.
func SemesterGT(v string) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSemester), v))
	})
}

// SemesterGTE applies the GTE predicate on the "semester" field.
func SemesterGTE(v string) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSemester), v))
	})
}

// SemesterLT applies the LT predicate on the "semester" field.
func SemesterLT(v string) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSemester), v))
	})
}

// SemesterLTE applies the LTE predicate on the "semester" field.
func SemesterLTE(v string) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSemester), v))
	})
}

// SemesterContains applies the Contains predicate on the "semester" field.
func SemesterContains(v string) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSemester), v))
	})
}

// SemesterHasPrefix applies the HasPrefix predicate on the "semester" field.
func SemesterHasPrefix(v string) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSemester), v))
	})
}

// SemesterHasSuffix applies the HasSuffix predicate on the "semester" field.
func SemesterHasSuffix(v string) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSemester), v))
	})
}

// SemesterEqualFold applies the EqualFold predicate on the "semester" field.
func SemesterEqualFold(v string) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSemester), v))
	})
}

// SemesterContainsFold applies the ContainsFold predicate on the "semester" field.
func SemesterContainsFold(v string) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSemester), v))
	})
}

// HasActivities applies the HasEdge predicate on the "activities" edge.
func HasActivities() predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActivitiesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ActivitiesTable, ActivitiesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActivitiesWith applies the HasEdge predicate on the "activities" edge with a given conditions (other predicates).
func HasActivitiesWith(preds ...predicate.Activities) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActivitiesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ActivitiesTable, ActivitiesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.AcademicYear) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.AcademicYear) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AcademicYear) predicate.AcademicYear {
	return predicate.AcademicYear(func(s *sql.Selector) {
		p(s.Not())
	})
}
