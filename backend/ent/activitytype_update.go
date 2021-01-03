// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/OMENX/app/ent/activities"
	"github.com/OMENX/app/ent/activitytype"
	"github.com/OMENX/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ActivityTypeUpdate is the builder for updating ActivityType entities.
type ActivityTypeUpdate struct {
	config
	hooks      []Hook
	mutation   *ActivityTypeMutation
	predicates []predicate.ActivityType
}

// Where adds a new predicate for the builder.
func (atu *ActivityTypeUpdate) Where(ps ...predicate.ActivityType) *ActivityTypeUpdate {
	atu.predicates = append(atu.predicates, ps...)
	return atu
}

// SetName sets the name field.
func (atu *ActivityTypeUpdate) SetName(s string) *ActivityTypeUpdate {
	atu.mutation.SetName(s)
	return atu
}

// AddActivityIDs adds the activities edge to Activities by ids.
func (atu *ActivityTypeUpdate) AddActivityIDs(ids ...int) *ActivityTypeUpdate {
	atu.mutation.AddActivityIDs(ids...)
	return atu
}

// AddActivities adds the activities edges to Activities.
func (atu *ActivityTypeUpdate) AddActivities(a ...*Activities) *ActivityTypeUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atu.AddActivityIDs(ids...)
}

// Mutation returns the ActivityTypeMutation object of the builder.
func (atu *ActivityTypeUpdate) Mutation() *ActivityTypeMutation {
	return atu.mutation
}

// RemoveActivityIDs removes the activities edge to Activities by ids.
func (atu *ActivityTypeUpdate) RemoveActivityIDs(ids ...int) *ActivityTypeUpdate {
	atu.mutation.RemoveActivityIDs(ids...)
	return atu
}

// RemoveActivities removes activities edges to Activities.
func (atu *ActivityTypeUpdate) RemoveActivities(a ...*Activities) *ActivityTypeUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atu.RemoveActivityIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (atu *ActivityTypeUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := atu.mutation.Name(); ok {
		if err := activitytype.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(atu.hooks) == 0 {
		affected, err = atu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActivityTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			atu.mutation = mutation
			affected, err = atu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atu.hooks) - 1; i >= 0; i-- {
			mut = atu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (atu *ActivityTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *ActivityTypeUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *ActivityTypeUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (atu *ActivityTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   activitytype.Table,
			Columns: activitytype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: activitytype.FieldID,
			},
		},
	}
	if ps := atu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: activitytype.FieldName,
		})
	}
	if nodes := atu.mutation.RemovedActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.ActivitiesTable,
			Columns: []string{activitytype.ActivitiesColumn},
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
	if nodes := atu.mutation.ActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.ActivitiesTable,
			Columns: []string{activitytype.ActivitiesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{activitytype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ActivityTypeUpdateOne is the builder for updating a single ActivityType entity.
type ActivityTypeUpdateOne struct {
	config
	hooks    []Hook
	mutation *ActivityTypeMutation
}

// SetName sets the name field.
func (atuo *ActivityTypeUpdateOne) SetName(s string) *ActivityTypeUpdateOne {
	atuo.mutation.SetName(s)
	return atuo
}

// AddActivityIDs adds the activities edge to Activities by ids.
func (atuo *ActivityTypeUpdateOne) AddActivityIDs(ids ...int) *ActivityTypeUpdateOne {
	atuo.mutation.AddActivityIDs(ids...)
	return atuo
}

// AddActivities adds the activities edges to Activities.
func (atuo *ActivityTypeUpdateOne) AddActivities(a ...*Activities) *ActivityTypeUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atuo.AddActivityIDs(ids...)
}

// Mutation returns the ActivityTypeMutation object of the builder.
func (atuo *ActivityTypeUpdateOne) Mutation() *ActivityTypeMutation {
	return atuo.mutation
}

// RemoveActivityIDs removes the activities edge to Activities by ids.
func (atuo *ActivityTypeUpdateOne) RemoveActivityIDs(ids ...int) *ActivityTypeUpdateOne {
	atuo.mutation.RemoveActivityIDs(ids...)
	return atuo
}

// RemoveActivities removes activities edges to Activities.
func (atuo *ActivityTypeUpdateOne) RemoveActivities(a ...*Activities) *ActivityTypeUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atuo.RemoveActivityIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (atuo *ActivityTypeUpdateOne) Save(ctx context.Context) (*ActivityType, error) {
	if v, ok := atuo.mutation.Name(); ok {
		if err := activitytype.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}

	var (
		err  error
		node *ActivityType
	)
	if len(atuo.hooks) == 0 {
		node, err = atuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActivityTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			atuo.mutation = mutation
			node, err = atuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(atuo.hooks) - 1; i >= 0; i-- {
			mut = atuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *ActivityTypeUpdateOne) SaveX(ctx context.Context) *ActivityType {
	at, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return at
}

// Exec executes the query on the entity.
func (atuo *ActivityTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *ActivityTypeUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (atuo *ActivityTypeUpdateOne) sqlSave(ctx context.Context) (at *ActivityType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   activitytype.Table,
			Columns: activitytype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: activitytype.FieldID,
			},
		},
	}
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ActivityType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := atuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: activitytype.FieldName,
		})
	}
	if nodes := atuo.mutation.RemovedActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.ActivitiesTable,
			Columns: []string{activitytype.ActivitiesColumn},
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
	if nodes := atuo.mutation.ActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activitytype.ActivitiesTable,
			Columns: []string{activitytype.ActivitiesColumn},
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
	at = &ActivityType{config: atuo.config}
	_spec.Assign = at.assignValues
	_spec.ScanValues = at.scanValues()
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{activitytype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return at, nil
}
