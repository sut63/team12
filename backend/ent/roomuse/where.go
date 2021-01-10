// Code generated by entc, DO NOT EDIT.

package roomuse

import (
	"time"

	"github.com/OMENX/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AddedTime applies equality check predicate on the "added_time" field. It's identical to AddedTimeEQ.
func AddedTime(v time.Time) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddedTime), v))
	})
}

// AddedTimeEQ applies the EQ predicate on the "added_time" field.
func AddedTimeEQ(v time.Time) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddedTime), v))
	})
}

// AddedTimeNEQ applies the NEQ predicate on the "added_time" field.
func AddedTimeNEQ(v time.Time) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddedTime), v))
	})
}

// AddedTimeIn applies the In predicate on the "added_time" field.
func AddedTimeIn(vs ...time.Time) predicate.Roomuse {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomuse(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddedTime), v...))
	})
}

// AddedTimeNotIn applies the NotIn predicate on the "added_time" field.
func AddedTimeNotIn(vs ...time.Time) predicate.Roomuse {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomuse(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddedTime), v...))
	})
}

// AddedTimeGT applies the GT predicate on the "added_time" field.
func AddedTimeGT(v time.Time) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddedTime), v))
	})
}

// AddedTimeGTE applies the GTE predicate on the "added_time" field.
func AddedTimeGTE(v time.Time) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddedTime), v))
	})
}

// AddedTimeLT applies the LT predicate on the "added_time" field.
func AddedTimeLT(v time.Time) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddedTime), v))
	})
}

// AddedTimeLTE applies the LTE predicate on the "added_time" field.
func AddedTimeLTE(v time.Time) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddedTime), v))
	})
}

// HasRooms applies the HasEdge predicate on the "rooms" edge.
func HasRooms() predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomsTable, RoomsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomsWith applies the HasEdge predicate on the "rooms" edge with a given conditions (other predicates).
func HasRoomsWith(preds ...predicate.Room) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomsTable, RoomsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPurposes applies the HasEdge predicate on the "purposes" edge.
func HasPurposes() predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PurposesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PurposesTable, PurposesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPurposesWith applies the HasEdge predicate on the "purposes" edge with a given conditions (other predicates).
func HasPurposesWith(preds ...predicate.Roompurpose) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PurposesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PurposesTable, PurposesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Roomuse) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Roomuse) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
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
func Not(p predicate.Roomuse) predicate.Roomuse {
	return predicate.Roomuse(func(s *sql.Selector) {
		p(s.Not())
	})
}