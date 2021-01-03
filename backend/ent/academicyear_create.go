// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/OMENX/app/ent/academicyear"
	"github.com/OMENX/app/ent/activities"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// AcademicYearCreate is the builder for creating a AcademicYear entity.
type AcademicYearCreate struct {
	config
	mutation *AcademicYearMutation
	hooks    []Hook
}

// SetSemester sets the semester field.
func (ayc *AcademicYearCreate) SetSemester(s string) *AcademicYearCreate {
	ayc.mutation.SetSemester(s)
	return ayc
}

// AddActivityIDs adds the activities edge to Activities by ids.
func (ayc *AcademicYearCreate) AddActivityIDs(ids ...int) *AcademicYearCreate {
	ayc.mutation.AddActivityIDs(ids...)
	return ayc
}

// AddActivities adds the activities edges to Activities.
func (ayc *AcademicYearCreate) AddActivities(a ...*Activities) *AcademicYearCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ayc.AddActivityIDs(ids...)
}

// Mutation returns the AcademicYearMutation object of the builder.
func (ayc *AcademicYearCreate) Mutation() *AcademicYearMutation {
	return ayc.mutation
}

// Save creates the AcademicYear in the database.
func (ayc *AcademicYearCreate) Save(ctx context.Context) (*AcademicYear, error) {
	if _, ok := ayc.mutation.Semester(); !ok {
		return nil, &ValidationError{Name: "semester", err: errors.New("ent: missing required field \"semester\"")}
	}
	if v, ok := ayc.mutation.Semester(); ok {
		if err := academicyear.SemesterValidator(v); err != nil {
			return nil, &ValidationError{Name: "semester", err: fmt.Errorf("ent: validator failed for field \"semester\": %w", err)}
		}
	}
	var (
		err  error
		node *AcademicYear
	)
	if len(ayc.hooks) == 0 {
		node, err = ayc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AcademicYearMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ayc.mutation = mutation
			node, err = ayc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ayc.hooks) - 1; i >= 0; i-- {
			mut = ayc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ayc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ayc *AcademicYearCreate) SaveX(ctx context.Context) *AcademicYear {
	v, err := ayc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ayc *AcademicYearCreate) sqlSave(ctx context.Context) (*AcademicYear, error) {
	ay, _spec := ayc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ayc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ay.ID = int(id)
	return ay, nil
}

func (ayc *AcademicYearCreate) createSpec() (*AcademicYear, *sqlgraph.CreateSpec) {
	var (
		ay    = &AcademicYear{config: ayc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: academicyear.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: academicyear.FieldID,
			},
		}
	)
	if value, ok := ayc.mutation.Semester(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: academicyear.FieldSemester,
		})
		ay.Semester = value
	}
	if nodes := ayc.mutation.ActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   academicyear.ActivitiesTable,
			Columns: []string{academicyear.ActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: activities.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return ay, _spec
}
