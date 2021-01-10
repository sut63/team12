// Code generated by entc, DO NOT EDIT.

package club

import (
	"time"

	"github.com/OMENX/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Purpose applies equality check predicate on the "purpose" field. It's identical to PurposeEQ.
func Purpose(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPurpose), v))
	})
}

// Foundingdate applies equality check predicate on the "foundingdate" field. It's identical to FoundingdateEQ.
func Foundingdate(v time.Time) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFoundingdate), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Club {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Club(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Club {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Club(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// PurposeEQ applies the EQ predicate on the "purpose" field.
func PurposeEQ(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPurpose), v))
	})
}

// PurposeNEQ applies the NEQ predicate on the "purpose" field.
func PurposeNEQ(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPurpose), v))
	})
}

// PurposeIn applies the In predicate on the "purpose" field.
func PurposeIn(vs ...string) predicate.Club {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Club(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPurpose), v...))
	})
}

// PurposeNotIn applies the NotIn predicate on the "purpose" field.
func PurposeNotIn(vs ...string) predicate.Club {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Club(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPurpose), v...))
	})
}

// PurposeGT applies the GT predicate on the "purpose" field.
func PurposeGT(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPurpose), v))
	})
}

// PurposeGTE applies the GTE predicate on the "purpose" field.
func PurposeGTE(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPurpose), v))
	})
}

// PurposeLT applies the LT predicate on the "purpose" field.
func PurposeLT(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPurpose), v))
	})
}

// PurposeLTE applies the LTE predicate on the "purpose" field.
func PurposeLTE(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPurpose), v))
	})
}

// PurposeContains applies the Contains predicate on the "purpose" field.
func PurposeContains(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPurpose), v))
	})
}

// PurposeHasPrefix applies the HasPrefix predicate on the "purpose" field.
func PurposeHasPrefix(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPurpose), v))
	})
}

// PurposeHasSuffix applies the HasSuffix predicate on the "purpose" field.
func PurposeHasSuffix(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPurpose), v))
	})
}

// PurposeEqualFold applies the EqualFold predicate on the "purpose" field.
func PurposeEqualFold(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPurpose), v))
	})
}

// PurposeContainsFold applies the ContainsFold predicate on the "purpose" field.
func PurposeContainsFold(v string) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPurpose), v))
	})
}

// FoundingdateEQ applies the EQ predicate on the "foundingdate" field.
func FoundingdateEQ(v time.Time) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFoundingdate), v))
	})
}

// FoundingdateNEQ applies the NEQ predicate on the "foundingdate" field.
func FoundingdateNEQ(v time.Time) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFoundingdate), v))
	})
}

// FoundingdateIn applies the In predicate on the "foundingdate" field.
func FoundingdateIn(vs ...time.Time) predicate.Club {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Club(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFoundingdate), v...))
	})
}

// FoundingdateNotIn applies the NotIn predicate on the "foundingdate" field.
func FoundingdateNotIn(vs ...time.Time) predicate.Club {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Club(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFoundingdate), v...))
	})
}

// FoundingdateGT applies the GT predicate on the "foundingdate" field.
func FoundingdateGT(v time.Time) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFoundingdate), v))
	})
}

// FoundingdateGTE applies the GTE predicate on the "foundingdate" field.
func FoundingdateGTE(v time.Time) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFoundingdate), v))
	})
}

// FoundingdateLT applies the LT predicate on the "foundingdate" field.
func FoundingdateLT(v time.Time) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFoundingdate), v))
	})
}

// FoundingdateLTE applies the LTE predicate on the "foundingdate" field.
func FoundingdateLTE(v time.Time) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFoundingdate), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClubtype applies the HasEdge predicate on the "clubtype" edge.
func HasClubtype() predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClubtypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClubtypeTable, ClubtypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClubtypeWith applies the HasEdge predicate on the "clubtype" edge with a given conditions (other predicates).
func HasClubtypeWith(preds ...predicate.ClubType) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClubtypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClubtypeTable, ClubtypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClubbranch applies the HasEdge predicate on the "clubbranch" edge.
func HasClubbranch() predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClubbranchTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClubbranchTable, ClubbranchColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClubbranchWith applies the HasEdge predicate on the "clubbranch" edge with a given conditions (other predicates).
func HasClubbranchWith(preds ...predicate.ClubBranch) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClubbranchInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClubbranchTable, ClubbranchColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClubapplication applies the HasEdge predicate on the "clubapplication" edge.
func HasClubapplication() predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClubapplicationTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClubapplicationTable, ClubapplicationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClubapplicationWith applies the HasEdge predicate on the "clubapplication" edge with a given conditions (other predicates).
func HasClubapplicationWith(preds ...predicate.Clubapplication) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
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

// HasClubToComplaint applies the HasEdge predicate on the "ClubToComplaint" edge.
func HasClubToComplaint() predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClubToComplaintTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClubToComplaintTable, ClubToComplaintColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClubToComplaintWith applies the HasEdge predicate on the "ClubToComplaint" edge with a given conditions (other predicates).
func HasClubToComplaintWith(preds ...predicate.Complaint) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClubToComplaintInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClubToComplaintTable, ClubToComplaintColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasActivities applies the HasEdge predicate on the "activities" edge.
func HasActivities() predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActivitiesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ActivitiesTable, ActivitiesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActivitiesWith applies the HasEdge predicate on the "activities" edge with a given conditions (other predicates).
func HasActivitiesWith(preds ...predicate.Activities) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
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
func And(predicates ...predicate.Club) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Club) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
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
func Not(p predicate.Club) predicate.Club {
	return predicate.Club(func(s *sql.Selector) {
		p(s.Not())
	})
}