// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/OMENX/app/ent/clubapplication"
	"github.com/OMENX/app/ent/clubappstatus"
	"github.com/OMENX/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ClubappStatusUpdate is the builder for updating ClubappStatus entities.
type ClubappStatusUpdate struct {
	config
	hooks      []Hook
	mutation   *ClubappStatusMutation
	predicates []predicate.ClubappStatus
}

// Where adds a new predicate for the builder.
func (csu *ClubappStatusUpdate) Where(ps ...predicate.ClubappStatus) *ClubappStatusUpdate {
	csu.predicates = append(csu.predicates, ps...)
	return csu
}

// SetClubstatus sets the clubstatus field.
func (csu *ClubappStatusUpdate) SetClubstatus(s string) *ClubappStatusUpdate {
	csu.mutation.SetClubstatus(s)
	return csu
}

// AddClubapplicationIDs adds the clubapplication edge to Clubapplication by ids.
func (csu *ClubappStatusUpdate) AddClubapplicationIDs(ids ...int) *ClubappStatusUpdate {
	csu.mutation.AddClubapplicationIDs(ids...)
	return csu
}

// AddClubapplication adds the clubapplication edges to Clubapplication.
func (csu *ClubappStatusUpdate) AddClubapplication(c ...*Clubapplication) *ClubappStatusUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return csu.AddClubapplicationIDs(ids...)
}

// Mutation returns the ClubappStatusMutation object of the builder.
func (csu *ClubappStatusUpdate) Mutation() *ClubappStatusMutation {
	return csu.mutation
}

// RemoveClubapplicationIDs removes the clubapplication edge to Clubapplication by ids.
func (csu *ClubappStatusUpdate) RemoveClubapplicationIDs(ids ...int) *ClubappStatusUpdate {
	csu.mutation.RemoveClubapplicationIDs(ids...)
	return csu
}

// RemoveClubapplication removes clubapplication edges to Clubapplication.
func (csu *ClubappStatusUpdate) RemoveClubapplication(c ...*Clubapplication) *ClubappStatusUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return csu.RemoveClubapplicationIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (csu *ClubappStatusUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(csu.hooks) == 0 {
		affected, err = csu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClubappStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			csu.mutation = mutation
			affected, err = csu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(csu.hooks) - 1; i >= 0; i-- {
			mut = csu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (csu *ClubappStatusUpdate) SaveX(ctx context.Context) int {
	affected, err := csu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (csu *ClubappStatusUpdate) Exec(ctx context.Context) error {
	_, err := csu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csu *ClubappStatusUpdate) ExecX(ctx context.Context) {
	if err := csu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (csu *ClubappStatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clubappstatus.Table,
			Columns: clubappstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clubappstatus.FieldID,
			},
		},
	}
	if ps := csu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csu.mutation.Clubstatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clubappstatus.FieldClubstatus,
		})
	}
	if nodes := csu.mutation.RemovedClubapplicationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   clubappstatus.ClubapplicationTable,
			Columns: []string{clubappstatus.ClubapplicationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubapplication.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csu.mutation.ClubapplicationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   clubappstatus.ClubapplicationTable,
			Columns: []string{clubappstatus.ClubapplicationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubapplication.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, csu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clubappstatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ClubappStatusUpdateOne is the builder for updating a single ClubappStatus entity.
type ClubappStatusUpdateOne struct {
	config
	hooks    []Hook
	mutation *ClubappStatusMutation
}

// SetClubstatus sets the clubstatus field.
func (csuo *ClubappStatusUpdateOne) SetClubstatus(s string) *ClubappStatusUpdateOne {
	csuo.mutation.SetClubstatus(s)
	return csuo
}

// AddClubapplicationIDs adds the clubapplication edge to Clubapplication by ids.
func (csuo *ClubappStatusUpdateOne) AddClubapplicationIDs(ids ...int) *ClubappStatusUpdateOne {
	csuo.mutation.AddClubapplicationIDs(ids...)
	return csuo
}

// AddClubapplication adds the clubapplication edges to Clubapplication.
func (csuo *ClubappStatusUpdateOne) AddClubapplication(c ...*Clubapplication) *ClubappStatusUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return csuo.AddClubapplicationIDs(ids...)
}

// Mutation returns the ClubappStatusMutation object of the builder.
func (csuo *ClubappStatusUpdateOne) Mutation() *ClubappStatusMutation {
	return csuo.mutation
}

// RemoveClubapplicationIDs removes the clubapplication edge to Clubapplication by ids.
func (csuo *ClubappStatusUpdateOne) RemoveClubapplicationIDs(ids ...int) *ClubappStatusUpdateOne {
	csuo.mutation.RemoveClubapplicationIDs(ids...)
	return csuo
}

// RemoveClubapplication removes clubapplication edges to Clubapplication.
func (csuo *ClubappStatusUpdateOne) RemoveClubapplication(c ...*Clubapplication) *ClubappStatusUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return csuo.RemoveClubapplicationIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (csuo *ClubappStatusUpdateOne) Save(ctx context.Context) (*ClubappStatus, error) {

	var (
		err  error
		node *ClubappStatus
	)
	if len(csuo.hooks) == 0 {
		node, err = csuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClubappStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			csuo.mutation = mutation
			node, err = csuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(csuo.hooks) - 1; i >= 0; i-- {
			mut = csuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (csuo *ClubappStatusUpdateOne) SaveX(ctx context.Context) *ClubappStatus {
	cs, err := csuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return cs
}

// Exec executes the query on the entity.
func (csuo *ClubappStatusUpdateOne) Exec(ctx context.Context) error {
	_, err := csuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csuo *ClubappStatusUpdateOne) ExecX(ctx context.Context) {
	if err := csuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (csuo *ClubappStatusUpdateOne) sqlSave(ctx context.Context) (cs *ClubappStatus, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clubappstatus.Table,
			Columns: clubappstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clubappstatus.FieldID,
			},
		},
	}
	id, ok := csuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ClubappStatus.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := csuo.mutation.Clubstatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clubappstatus.FieldClubstatus,
		})
	}
	if nodes := csuo.mutation.RemovedClubapplicationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   clubappstatus.ClubapplicationTable,
			Columns: []string{clubappstatus.ClubapplicationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubapplication.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := csuo.mutation.ClubapplicationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   clubappstatus.ClubapplicationTable,
			Columns: []string{clubappstatus.ClubapplicationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubapplication.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	cs = &ClubappStatus{config: csuo.config}
	_spec.Assign = cs.assignValues
	_spec.ScanValues = cs.scanValues()
	if err = sqlgraph.UpdateNode(ctx, csuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clubappstatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return cs, nil
}
