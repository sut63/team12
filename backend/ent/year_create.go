// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/OMENX/app/ent/user"
	"github.com/OMENX/app/ent/year"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// YearCreate is the builder for creating a Year entity.
type YearCreate struct {
	config
	mutation *YearMutation
	hooks    []Hook
}

// SetYear sets the year field.
func (yc *YearCreate) SetYear(i int) *YearCreate {
	yc.mutation.SetYear(i)
	return yc
}

// AddUserIDs adds the users edge to User by ids.
func (yc *YearCreate) AddUserIDs(ids ...int) *YearCreate {
	yc.mutation.AddUserIDs(ids...)
	return yc
}

// AddUsers adds the users edges to User.
func (yc *YearCreate) AddUsers(u ...*User) *YearCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return yc.AddUserIDs(ids...)
}

// Mutation returns the YearMutation object of the builder.
func (yc *YearCreate) Mutation() *YearMutation {
	return yc.mutation
}

// Save creates the Year in the database.
func (yc *YearCreate) Save(ctx context.Context) (*Year, error) {
	if _, ok := yc.mutation.Year(); !ok {
		return nil, &ValidationError{Name: "year", err: errors.New("ent: missing required field \"year\"")}
	}
	var (
		err  error
		node *Year
	)
	if len(yc.hooks) == 0 {
		node, err = yc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*YearMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			yc.mutation = mutation
			node, err = yc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(yc.hooks) - 1; i >= 0; i-- {
			mut = yc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, yc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (yc *YearCreate) SaveX(ctx context.Context) *Year {
	v, err := yc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (yc *YearCreate) sqlSave(ctx context.Context) (*Year, error) {
	y, _spec := yc.createSpec()
	if err := sqlgraph.CreateNode(ctx, yc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	y.ID = int(id)
	return y, nil
}

func (yc *YearCreate) createSpec() (*Year, *sqlgraph.CreateSpec) {
	var (
		y     = &Year{config: yc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: year.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: year.FieldID,
			},
		}
	)
	if value, ok := yc.mutation.Year(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: year.FieldYear,
		})
		y.Year = value
	}
	if nodes := yc.mutation.UsersIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return y, _spec
}
