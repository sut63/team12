// Code generated by entc, DO NOT EDIT.

package clubappstatus

import (
	"github.com/OMENX/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ApplyStatus applies equality check predicate on the "apply_status" field. It's identical to ApplyStatusEQ.
func ApplyStatus(v string) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApplyStatus), v))
	})
}

// ApplyStatusEQ applies the EQ predicate on the "apply_status" field.
func ApplyStatusEQ(v string) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApplyStatus), v))
	})
}

// ApplyStatusNEQ applies the NEQ predicate on the "apply_status" field.
func ApplyStatusNEQ(v string) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldApplyStatus), v))
	})
}

// ApplyStatusIn applies the In predicate on the "apply_status" field.
func ApplyStatusIn(vs ...string) predicate.ClubappStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClubappStatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldApplyStatus), v...))
	})
}

// ApplyStatusNotIn applies the NotIn predicate on the "apply_status" field.
func ApplyStatusNotIn(vs ...string) predicate.ClubappStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClubappStatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldApplyStatus), v...))
	})
}

// ApplyStatusGT applies the GT predicate on the "apply_status" field.
func ApplyStatusGT(v string) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldApplyStatus), v))
	})
}

// ApplyStatusGTE applies the GTE predicate on the "apply_status" field.
func ApplyStatusGTE(v string) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldApplyStatus), v))
	})
}

// ApplyStatusLT applies the LT predicate on the "apply_status" field.
func ApplyStatusLT(v string) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldApplyStatus), v))
	})
}

// ApplyStatusLTE applies the LTE predicate on the "apply_status" field.
func ApplyStatusLTE(v string) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldApplyStatus), v))
	})
}

// ApplyStatusContains applies the Contains predicate on the "apply_status" field.
func ApplyStatusContains(v string) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldApplyStatus), v))
	})
}

// ApplyStatusHasPrefix applies the HasPrefix predicate on the "apply_status" field.
func ApplyStatusHasPrefix(v string) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldApplyStatus), v))
	})
}

// ApplyStatusHasSuffix applies the HasSuffix predicate on the "apply_status" field.
func ApplyStatusHasSuffix(v string) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldApplyStatus), v))
	})
}

// ApplyStatusEqualFold applies the EqualFold predicate on the "apply_status" field.
func ApplyStatusEqualFold(v string) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldApplyStatus), v))
	})
}

// ApplyStatusContainsFold applies the ContainsFold predicate on the "apply_status" field.
func ApplyStatusContainsFold(v string) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldApplyStatus), v))
	})
}

// HasClubapplication applies the HasEdge predicate on the "clubapplication" edge.
func HasClubapplication() predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClubapplicationTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClubapplicationTable, ClubapplicationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClubapplicationWith applies the HasEdge predicate on the "clubapplication" edge with a given conditions (other predicates).
func HasClubapplicationWith(preds ...predicate.Clubapplication) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClubapplicationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClubapplicationTable, ClubapplicationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.ClubappStatus) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.ClubappStatus) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
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
func Not(p predicate.ClubappStatus) predicate.ClubappStatus {
	return predicate.ClubappStatus(func(s *sql.Selector) {
		p(s.Not())
	})
}
