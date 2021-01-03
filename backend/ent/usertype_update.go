// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/OMENX/app/ent/predicate"
	"github.com/OMENX/app/ent/user"
	"github.com/OMENX/app/ent/usertype"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UsertypeUpdate is the builder for updating Usertype entities.
type UsertypeUpdate struct {
	config
	hooks      []Hook
	mutation   *UsertypeMutation
	predicates []predicate.Usertype
}

// Where adds a new predicate for the builder.
func (uu *UsertypeUpdate) Where(ps ...predicate.Usertype) *UsertypeUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetName sets the name field.
func (uu *UsertypeUpdate) SetName(s string) *UsertypeUpdate {
	uu.mutation.SetName(s)
	return uu
}

// AddUserIDs adds the user edge to User by ids.
func (uu *UsertypeUpdate) AddUserIDs(ids ...int) *UsertypeUpdate {
	uu.mutation.AddUserIDs(ids...)
	return uu
}

// AddUser adds the user edges to User.
func (uu *UsertypeUpdate) AddUser(u ...*User) *UsertypeUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddUserIDs(ids...)
}

// Mutation returns the UsertypeMutation object of the builder.
func (uu *UsertypeUpdate) Mutation() *UsertypeMutation {
	return uu.mutation
}

// RemoveUserIDs removes the user edge to User by ids.
func (uu *UsertypeUpdate) RemoveUserIDs(ids ...int) *UsertypeUpdate {
	uu.mutation.RemoveUserIDs(ids...)
	return uu
}

// RemoveUser removes user edges to User.
func (uu *UsertypeUpdate) RemoveUser(u ...*User) *UsertypeUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UsertypeUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := uu.mutation.Name(); ok {
		if err := usertype.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UsertypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UsertypeUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UsertypeUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UsertypeUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UsertypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usertype.Table,
			Columns: usertype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usertype.FieldID,
			},
		},
	}
	if ps := uu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usertype.FieldName,
		})
	}
	if nodes := uu.mutation.RemovedUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usertype.UserTable,
			Columns: []string{usertype.UserColumn},
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
	if nodes := uu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usertype.UserTable,
			Columns: []string{usertype.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usertype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UsertypeUpdateOne is the builder for updating a single Usertype entity.
type UsertypeUpdateOne struct {
	config
	hooks    []Hook
	mutation *UsertypeMutation
}

// SetName sets the name field.
func (uuo *UsertypeUpdateOne) SetName(s string) *UsertypeUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// AddUserIDs adds the user edge to User by ids.
func (uuo *UsertypeUpdateOne) AddUserIDs(ids ...int) *UsertypeUpdateOne {
	uuo.mutation.AddUserIDs(ids...)
	return uuo
}

// AddUser adds the user edges to User.
func (uuo *UsertypeUpdateOne) AddUser(u ...*User) *UsertypeUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddUserIDs(ids...)
}

// Mutation returns the UsertypeMutation object of the builder.
func (uuo *UsertypeUpdateOne) Mutation() *UsertypeMutation {
	return uuo.mutation
}

// RemoveUserIDs removes the user edge to User by ids.
func (uuo *UsertypeUpdateOne) RemoveUserIDs(ids ...int) *UsertypeUpdateOne {
	uuo.mutation.RemoveUserIDs(ids...)
	return uuo
}

// RemoveUser removes user edges to User.
func (uuo *UsertypeUpdateOne) RemoveUser(u ...*User) *UsertypeUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveUserIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (uuo *UsertypeUpdateOne) Save(ctx context.Context) (*Usertype, error) {
	if v, ok := uuo.mutation.Name(); ok {
		if err := usertype.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}

	var (
		err  error
		node *Usertype
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UsertypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UsertypeUpdateOne) SaveX(ctx context.Context) *Usertype {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UsertypeUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UsertypeUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UsertypeUpdateOne) sqlSave(ctx context.Context) (u *Usertype, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usertype.Table,
			Columns: usertype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usertype.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Usertype.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := uuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usertype.FieldName,
		})
	}
	if nodes := uuo.mutation.RemovedUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usertype.UserTable,
			Columns: []string{usertype.UserColumn},
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
	if nodes := uuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usertype.UserTable,
			Columns: []string{usertype.UserColumn},
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
	u = &Usertype{config: uuo.config}
	_spec.Assign = u.assignValues
	_spec.ScanValues = u.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usertype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return u, nil
}
