// Code generated by entc, DO NOT EDIT.

package complaint

import (
	"github.com/OMENX/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Info applies equality check predicate on the "info" field. It's identical to InfoEQ.
func Info(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInfo), v))
	})
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// InfoEQ applies the EQ predicate on the "info" field.
func InfoEQ(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInfo), v))
	})
}

// InfoNEQ applies the NEQ predicate on the "info" field.
func InfoNEQ(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInfo), v))
	})
}

// InfoIn applies the In predicate on the "info" field.
func InfoIn(vs ...string) predicate.Complaint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Complaint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInfo), v...))
	})
}

// InfoNotIn applies the NotIn predicate on the "info" field.
func InfoNotIn(vs ...string) predicate.Complaint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Complaint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInfo), v...))
	})
}

// InfoGT applies the GT predicate on the "info" field.
func InfoGT(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInfo), v))
	})
}

// InfoGTE applies the GTE predicate on the "info" field.
func InfoGTE(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInfo), v))
	})
}

// InfoLT applies the LT predicate on the "info" field.
func InfoLT(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInfo), v))
	})
}

// InfoLTE applies the LTE predicate on the "info" field.
func InfoLTE(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInfo), v))
	})
}

// InfoContains applies the Contains predicate on the "info" field.
func InfoContains(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInfo), v))
	})
}

// InfoHasPrefix applies the HasPrefix predicate on the "info" field.
func InfoHasPrefix(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInfo), v))
	})
}

// InfoHasSuffix applies the HasSuffix predicate on the "info" field.
func InfoHasSuffix(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInfo), v))
	})
}

// InfoEqualFold applies the EqualFold predicate on the "info" field.
func InfoEqualFold(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInfo), v))
	})
}

// InfoContainsFold applies the ContainsFold predicate on the "info" field.
func InfoContainsFold(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInfo), v))
	})
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDate), v))
	})
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...string) predicate.Complaint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Complaint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDate), v...))
	})
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...string) predicate.Complaint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Complaint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDate), v...))
	})
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDate), v))
	})
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDate), v))
	})
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDate), v))
	})
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDate), v))
	})
}

// DateContains applies the Contains predicate on the "date" field.
func DateContains(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDate), v))
	})
}

// DateHasPrefix applies the HasPrefix predicate on the "date" field.
func DateHasPrefix(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDate), v))
	})
}

// DateHasSuffix applies the HasSuffix predicate on the "date" field.
func DateHasSuffix(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDate), v))
	})
}

// DateEqualFold applies the EqualFold predicate on the "date" field.
func DateEqualFold(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDate), v))
	})
}

// DateContainsFold applies the ContainsFold predicate on the "date" field.
func DateContainsFold(v string) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDate), v))
	})
}

// HasComplaintToUser applies the HasEdge predicate on the "ComplaintToUser" edge.
func HasComplaintToUser() predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ComplaintToUserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ComplaintToUserTable, ComplaintToUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasComplaintToUserWith applies the HasEdge predicate on the "ComplaintToUser" edge with a given conditions (other predicates).
func HasComplaintToUserWith(preds ...predicate.User) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ComplaintToUserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ComplaintToUserTable, ComplaintToUserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComplaintToComplaintType applies the HasEdge predicate on the "ComplaintToComplaintType" edge.
func HasComplaintToComplaintType() predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ComplaintToComplaintTypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ComplaintToComplaintTypeTable, ComplaintToComplaintTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasComplaintToComplaintTypeWith applies the HasEdge predicate on the "ComplaintToComplaintType" edge with a given conditions (other predicates).
func HasComplaintToComplaintTypeWith(preds ...predicate.ComplaintType) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ComplaintToComplaintTypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ComplaintToComplaintTypeTable, ComplaintToComplaintTypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComplaintToClub applies the HasEdge predicate on the "ComplaintToClub" edge.
func HasComplaintToClub() predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ComplaintToClubTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ComplaintToClubTable, ComplaintToClubColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasComplaintToClubWith applies the HasEdge predicate on the "ComplaintToClub" edge with a given conditions (other predicates).
func HasComplaintToClubWith(preds ...predicate.Club) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ComplaintToClubInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ComplaintToClubTable, ComplaintToClubColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Complaint) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Complaint) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
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
func Not(p predicate.Complaint) predicate.Complaint {
	return predicate.Complaint(func(s *sql.Selector) {
		p(s.Not())
	})
}
