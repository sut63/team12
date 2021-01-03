// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/OMENX/app/ent/academicyear"
	"github.com/OMENX/app/ent/activities"
	"github.com/OMENX/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// AcademicYearUpdate is the builder for updating AcademicYear entities.
type AcademicYearUpdate struct {
	config
	hooks      []Hook
	mutation   *AcademicYearMutation
	predicates []predicate.AcademicYear
}

// Where adds a new predicate for the builder.
func (ayu *AcademicYearUpdate) Where(ps ...predicate.AcademicYear) *AcademicYearUpdate {
	ayu.predicates = append(ayu.predicates, ps...)
	return ayu
}

// SetSemester sets the semester field.
func (ayu *AcademicYearUpdate) SetSemester(s string) *AcademicYearUpdate {
	ayu.mutation.SetSemester(s)
	return ayu
}

// AddActivityIDs adds the activities edge to Activities by ids.
func (ayu *AcademicYearUpdate) AddActivityIDs(ids ...int) *AcademicYearUpdate {
	ayu.mutation.AddActivityIDs(ids...)
	return ayu
}

// AddActivities adds the activities edges to Activities.
func (ayu *AcademicYearUpdate) AddActivities(a ...*Activities) *AcademicYearUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ayu.AddActivityIDs(ids...)
}

// Mutation returns the AcademicYearMutation object of the builder.
func (ayu *AcademicYearUpdate) Mutation() *AcademicYearMutation {
	return ayu.mutation
}

// RemoveActivityIDs removes the activities edge to Activities by ids.
func (ayu *AcademicYearUpdate) RemoveActivityIDs(ids ...int) *AcademicYearUpdate {
	ayu.mutation.RemoveActivityIDs(ids...)
	return ayu
}

// RemoveActivities removes activities edges to Activities.
func (ayu *AcademicYearUpdate) RemoveActivities(a ...*Activities) *AcademicYearUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ayu.RemoveActivityIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ayu *AcademicYearUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := ayu.mutation.Semester(); ok {
		if err := academicyear.SemesterValidator(v); err != nil {
			return 0, &ValidationError{Name: "semester", err: fmt.Errorf("ent: validator failed for field \"semester\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(ayu.hooks) == 0 {
		affected, err = ayu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AcademicYearMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ayu.mutation = mutation
			affected, err = ayu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ayu.hooks) - 1; i >= 0; i-- {
			mut = ayu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ayu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ayu *AcademicYearUpdate) SaveX(ctx context.Context) int {
	affected, err := ayu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ayu *AcademicYearUpdate) Exec(ctx context.Context) error {
	_, err := ayu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ayu *AcademicYearUpdate) ExecX(ctx context.Context) {
	if err := ayu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ayu *AcademicYearUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   academicyear.Table,
			Columns: academicyear.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: academicyear.FieldID,
			},
		},
	}
	if ps := ayu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ayu.mutation.Semester(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: academicyear.FieldSemester,
		})
	}
	if nodes := ayu.mutation.RemovedActivitiesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ayu.mutation.ActivitiesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ayu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{academicyear.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AcademicYearUpdateOne is the builder for updating a single AcademicYear entity.
type AcademicYearUpdateOne struct {
	config
	hooks    []Hook
	mutation *AcademicYearMutation
}

// SetSemester sets the semester field.
func (ayuo *AcademicYearUpdateOne) SetSemester(s string) *AcademicYearUpdateOne {
	ayuo.mutation.SetSemester(s)
	return ayuo
}

// AddActivityIDs adds the activities edge to Activities by ids.
func (ayuo *AcademicYearUpdateOne) AddActivityIDs(ids ...int) *AcademicYearUpdateOne {
	ayuo.mutation.AddActivityIDs(ids...)
	return ayuo
}

// AddActivities adds the activities edges to Activities.
func (ayuo *AcademicYearUpdateOne) AddActivities(a ...*Activities) *AcademicYearUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ayuo.AddActivityIDs(ids...)
}

// Mutation returns the AcademicYearMutation object of the builder.
func (ayuo *AcademicYearUpdateOne) Mutation() *AcademicYearMutation {
	return ayuo.mutation
}

// RemoveActivityIDs removes the activities edge to Activities by ids.
func (ayuo *AcademicYearUpdateOne) RemoveActivityIDs(ids ...int) *AcademicYearUpdateOne {
	ayuo.mutation.RemoveActivityIDs(ids...)
	return ayuo
}

// RemoveActivities removes activities edges to Activities.
func (ayuo *AcademicYearUpdateOne) RemoveActivities(a ...*Activities) *AcademicYearUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ayuo.RemoveActivityIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (ayuo *AcademicYearUpdateOne) Save(ctx context.Context) (*AcademicYear, error) {
	if v, ok := ayuo.mutation.Semester(); ok {
		if err := academicyear.SemesterValidator(v); err != nil {
			return nil, &ValidationError{Name: "semester", err: fmt.Errorf("ent: validator failed for field \"semester\": %w", err)}
		}
	}

	var (
		err  error
		node *AcademicYear
	)
	if len(ayuo.hooks) == 0 {
		node, err = ayuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AcademicYearMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ayuo.mutation = mutation
			node, err = ayuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ayuo.hooks) - 1; i >= 0; i-- {
			mut = ayuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ayuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ayuo *AcademicYearUpdateOne) SaveX(ctx context.Context) *AcademicYear {
	ay, err := ayuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return ay
}

// Exec executes the query on the entity.
func (ayuo *AcademicYearUpdateOne) Exec(ctx context.Context) error {
	_, err := ayuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ayuo *AcademicYearUpdateOne) ExecX(ctx context.Context) {
	if err := ayuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ayuo *AcademicYearUpdateOne) sqlSave(ctx context.Context) (ay *AcademicYear, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   academicyear.Table,
			Columns: academicyear.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: academicyear.FieldID,
			},
		},
	}
	id, ok := ayuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing AcademicYear.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ayuo.mutation.Semester(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: academicyear.FieldSemester,
		})
	}
	if nodes := ayuo.mutation.RemovedActivitiesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ayuo.mutation.ActivitiesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	ay = &AcademicYear{config: ayuo.config}
	_spec.Assign = ay.assignValues
	_spec.ScanValues = ay.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ayuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{academicyear.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return ay, nil
}
