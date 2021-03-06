// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/OMENX/app/ent/predicate"
	"github.com/OMENX/app/ent/user"
	"github.com/OMENX/app/ent/year"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// YearUpdate is the builder for updating Year entities.
type YearUpdate struct {
	config
	hooks      []Hook
	mutation   *YearMutation
	predicates []predicate.Year
}

// Where adds a new predicate for the builder.
func (yu *YearUpdate) Where(ps ...predicate.Year) *YearUpdate {
	yu.predicates = append(yu.predicates, ps...)
	return yu
}

// SetYear sets the year field.
func (yu *YearUpdate) SetYear(i int) *YearUpdate {
	yu.mutation.ResetYear()
	yu.mutation.SetYear(i)
	return yu
}

// AddYear adds i to year.
func (yu *YearUpdate) AddYear(i int) *YearUpdate {
	yu.mutation.AddYear(i)
	return yu
}

// AddUserIDs adds the users edge to User by ids.
func (yu *YearUpdate) AddUserIDs(ids ...int) *YearUpdate {
	yu.mutation.AddUserIDs(ids...)
	return yu
}

// AddUsers adds the users edges to User.
func (yu *YearUpdate) AddUsers(u ...*User) *YearUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return yu.AddUserIDs(ids...)
}

// Mutation returns the YearMutation object of the builder.
func (yu *YearUpdate) Mutation() *YearMutation {
	return yu.mutation
}

// RemoveUserIDs removes the users edge to User by ids.
func (yu *YearUpdate) RemoveUserIDs(ids ...int) *YearUpdate {
	yu.mutation.RemoveUserIDs(ids...)
	return yu
}

// RemoveUsers removes users edges to User.
func (yu *YearUpdate) RemoveUsers(u ...*User) *YearUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return yu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (yu *YearUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(yu.hooks) == 0 {
		affected, err = yu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*YearMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			yu.mutation = mutation
			affected, err = yu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(yu.hooks) - 1; i >= 0; i-- {
			mut = yu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, yu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (yu *YearUpdate) SaveX(ctx context.Context) int {
	affected, err := yu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (yu *YearUpdate) Exec(ctx context.Context) error {
	_, err := yu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (yu *YearUpdate) ExecX(ctx context.Context) {
	if err := yu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (yu *YearUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   year.Table,
			Columns: year.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: year.FieldID,
			},
		},
	}
	if ps := yu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := yu.mutation.Year(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: year.FieldYear,
		})
	}
	if value, ok := yu.mutation.AddedYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: year.FieldYear,
		})
	}
	if nodes := yu.mutation.RemovedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.UsersTable,
			Columns: []string{year.UsersColumn},
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
	if nodes := yu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.UsersTable,
			Columns: []string{year.UsersColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, yu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{year.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// YearUpdateOne is the builder for updating a single Year entity.
type YearUpdateOne struct {
	config
	hooks    []Hook
	mutation *YearMutation
}

// SetYear sets the year field.
func (yuo *YearUpdateOne) SetYear(i int) *YearUpdateOne {
	yuo.mutation.ResetYear()
	yuo.mutation.SetYear(i)
	return yuo
}

// AddYear adds i to year.
func (yuo *YearUpdateOne) AddYear(i int) *YearUpdateOne {
	yuo.mutation.AddYear(i)
	return yuo
}

// AddUserIDs adds the users edge to User by ids.
func (yuo *YearUpdateOne) AddUserIDs(ids ...int) *YearUpdateOne {
	yuo.mutation.AddUserIDs(ids...)
	return yuo
}

// AddUsers adds the users edges to User.
func (yuo *YearUpdateOne) AddUsers(u ...*User) *YearUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return yuo.AddUserIDs(ids...)
}

// Mutation returns the YearMutation object of the builder.
func (yuo *YearUpdateOne) Mutation() *YearMutation {
	return yuo.mutation
}

// RemoveUserIDs removes the users edge to User by ids.
func (yuo *YearUpdateOne) RemoveUserIDs(ids ...int) *YearUpdateOne {
	yuo.mutation.RemoveUserIDs(ids...)
	return yuo
}

// RemoveUsers removes users edges to User.
func (yuo *YearUpdateOne) RemoveUsers(u ...*User) *YearUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return yuo.RemoveUserIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (yuo *YearUpdateOne) Save(ctx context.Context) (*Year, error) {

	var (
		err  error
		node *Year
	)
	if len(yuo.hooks) == 0 {
		node, err = yuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*YearMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			yuo.mutation = mutation
			node, err = yuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(yuo.hooks) - 1; i >= 0; i-- {
			mut = yuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, yuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (yuo *YearUpdateOne) SaveX(ctx context.Context) *Year {
	y, err := yuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return y
}

// Exec executes the query on the entity.
func (yuo *YearUpdateOne) Exec(ctx context.Context) error {
	_, err := yuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (yuo *YearUpdateOne) ExecX(ctx context.Context) {
	if err := yuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (yuo *YearUpdateOne) sqlSave(ctx context.Context) (y *Year, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   year.Table,
			Columns: year.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: year.FieldID,
			},
		},
	}
	id, ok := yuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Year.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := yuo.mutation.Year(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: year.FieldYear,
		})
	}
	if value, ok := yuo.mutation.AddedYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: year.FieldYear,
		})
	}
	if nodes := yuo.mutation.RemovedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.UsersTable,
			Columns: []string{year.UsersColumn},
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
	if nodes := yuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.UsersTable,
			Columns: []string{year.UsersColumn},
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
	y = &Year{config: yuo.config}
	_spec.Assign = y.assignValues
	_spec.ScanValues = y.scanValues()
	if err = sqlgraph.UpdateNode(ctx, yuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{year.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return y, nil
}
