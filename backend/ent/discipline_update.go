// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/OMENX/app/ent/discipline"
	"github.com/OMENX/app/ent/predicate"
	"github.com/OMENX/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DisciplineUpdate is the builder for updating Discipline entities.
type DisciplineUpdate struct {
	config
	hooks      []Hook
	mutation   *DisciplineMutation
	predicates []predicate.Discipline
}

// Where adds a new predicate for the builder.
func (du *DisciplineUpdate) Where(ps ...predicate.Discipline) *DisciplineUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetDiscipline sets the discipline field.
func (du *DisciplineUpdate) SetDiscipline(s string) *DisciplineUpdate {
	du.mutation.SetDiscipline(s)
	return du
}

// AddUserIDs adds the users edge to User by ids.
func (du *DisciplineUpdate) AddUserIDs(ids ...int) *DisciplineUpdate {
	du.mutation.AddUserIDs(ids...)
	return du
}

// AddUsers adds the users edges to User.
func (du *DisciplineUpdate) AddUsers(u ...*User) *DisciplineUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.AddUserIDs(ids...)
}

// Mutation returns the DisciplineMutation object of the builder.
func (du *DisciplineUpdate) Mutation() *DisciplineMutation {
	return du.mutation
}

// RemoveUserIDs removes the users edge to User by ids.
func (du *DisciplineUpdate) RemoveUserIDs(ids ...int) *DisciplineUpdate {
	du.mutation.RemoveUserIDs(ids...)
	return du
}

// RemoveUsers removes users edges to User.
func (du *DisciplineUpdate) RemoveUsers(u ...*User) *DisciplineUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DisciplineUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DisciplineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DisciplineUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DisciplineUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DisciplineUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DisciplineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   discipline.Table,
			Columns: discipline.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: discipline.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Discipline(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discipline.FieldDiscipline,
		})
	}
	if nodes := du.mutation.RemovedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   discipline.UsersTable,
			Columns: []string{discipline.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   discipline.UsersTable,
			Columns: []string{discipline.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{discipline.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DisciplineUpdateOne is the builder for updating a single Discipline entity.
type DisciplineUpdateOne struct {
	config
	hooks    []Hook
	mutation *DisciplineMutation
}

// SetDiscipline sets the discipline field.
func (duo *DisciplineUpdateOne) SetDiscipline(s string) *DisciplineUpdateOne {
	duo.mutation.SetDiscipline(s)
	return duo
}

// AddUserIDs adds the users edge to User by ids.
func (duo *DisciplineUpdateOne) AddUserIDs(ids ...int) *DisciplineUpdateOne {
	duo.mutation.AddUserIDs(ids...)
	return duo
}

// AddUsers adds the users edges to User.
func (duo *DisciplineUpdateOne) AddUsers(u ...*User) *DisciplineUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.AddUserIDs(ids...)
}

// Mutation returns the DisciplineMutation object of the builder.
func (duo *DisciplineUpdateOne) Mutation() *DisciplineMutation {
	return duo.mutation
}

// RemoveUserIDs removes the users edge to User by ids.
func (duo *DisciplineUpdateOne) RemoveUserIDs(ids ...int) *DisciplineUpdateOne {
	duo.mutation.RemoveUserIDs(ids...)
	return duo
}

// RemoveUsers removes users edges to User.
func (duo *DisciplineUpdateOne) RemoveUsers(u ...*User) *DisciplineUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.RemoveUserIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DisciplineUpdateOne) Save(ctx context.Context) (*Discipline, error) {

	var (
		err  error
		node *Discipline
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DisciplineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DisciplineUpdateOne) SaveX(ctx context.Context) *Discipline {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DisciplineUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DisciplineUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DisciplineUpdateOne) sqlSave(ctx context.Context) (d *Discipline, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   discipline.Table,
			Columns: discipline.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: discipline.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Discipline.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Discipline(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discipline.FieldDiscipline,
		})
	}
	if nodes := duo.mutation.RemovedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   discipline.UsersTable,
			Columns: []string{discipline.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   discipline.UsersTable,
			Columns: []string{discipline.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	d = &Discipline{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{discipline.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}
