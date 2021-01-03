// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/OMENX/app/ent/academicyear"
	"github.com/OMENX/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// AcademicYearDelete is the builder for deleting a AcademicYear entity.
type AcademicYearDelete struct {
	config
	hooks      []Hook
	mutation   *AcademicYearMutation
	predicates []predicate.AcademicYear
}

// Where adds a new predicate to the delete builder.
func (ayd *AcademicYearDelete) Where(ps ...predicate.AcademicYear) *AcademicYearDelete {
	ayd.predicates = append(ayd.predicates, ps...)
	return ayd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ayd *AcademicYearDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ayd.hooks) == 0 {
		affected, err = ayd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AcademicYearMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ayd.mutation = mutation
			affected, err = ayd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ayd.hooks) - 1; i >= 0; i-- {
			mut = ayd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ayd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ayd *AcademicYearDelete) ExecX(ctx context.Context) int {
	n, err := ayd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ayd *AcademicYearDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: academicyear.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: academicyear.FieldID,
			},
		},
	}
	if ps := ayd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ayd.driver, _spec)
}

// AcademicYearDeleteOne is the builder for deleting a single AcademicYear entity.
type AcademicYearDeleteOne struct {
	ayd *AcademicYearDelete
}

// Exec executes the deletion query.
func (aydo *AcademicYearDeleteOne) Exec(ctx context.Context) error {
	n, err := aydo.ayd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{academicyear.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aydo *AcademicYearDeleteOne) ExecX(ctx context.Context) {
	aydo.ayd.ExecX(ctx)
}
