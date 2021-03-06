// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/OMENX/app/ent/clubapplication"
	"github.com/OMENX/app/ent/clubappstatus"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ClubappStatusCreate is the builder for creating a ClubappStatus entity.
type ClubappStatusCreate struct {
	config
	mutation *ClubappStatusMutation
	hooks    []Hook
}

// SetClubstatus sets the clubstatus field.
func (csc *ClubappStatusCreate) SetClubstatus(s string) *ClubappStatusCreate {
	csc.mutation.SetClubstatus(s)
	return csc
}

// AddClubapplicationIDs adds the clubapplication edge to Clubapplication by ids.
func (csc *ClubappStatusCreate) AddClubapplicationIDs(ids ...int) *ClubappStatusCreate {
	csc.mutation.AddClubapplicationIDs(ids...)
	return csc
}

// AddClubapplication adds the clubapplication edges to Clubapplication.
func (csc *ClubappStatusCreate) AddClubapplication(c ...*Clubapplication) *ClubappStatusCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return csc.AddClubapplicationIDs(ids...)
}

// Mutation returns the ClubappStatusMutation object of the builder.
func (csc *ClubappStatusCreate) Mutation() *ClubappStatusMutation {
	return csc.mutation
}

// Save creates the ClubappStatus in the database.
func (csc *ClubappStatusCreate) Save(ctx context.Context) (*ClubappStatus, error) {
	if _, ok := csc.mutation.Clubstatus(); !ok {
		return nil, &ValidationError{Name: "clubstatus", err: errors.New("ent: missing required field \"clubstatus\"")}
	}
	var (
		err  error
		node *ClubappStatus
	)
	if len(csc.hooks) == 0 {
		node, err = csc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClubappStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			csc.mutation = mutation
			node, err = csc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(csc.hooks) - 1; i >= 0; i-- {
			mut = csc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (csc *ClubappStatusCreate) SaveX(ctx context.Context) *ClubappStatus {
	v, err := csc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (csc *ClubappStatusCreate) sqlSave(ctx context.Context) (*ClubappStatus, error) {
	cs, _spec := csc.createSpec()
	if err := sqlgraph.CreateNode(ctx, csc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	cs.ID = int(id)
	return cs, nil
}

func (csc *ClubappStatusCreate) createSpec() (*ClubappStatus, *sqlgraph.CreateSpec) {
	var (
		cs    = &ClubappStatus{config: csc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: clubappstatus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clubappstatus.FieldID,
			},
		}
	)
	if value, ok := csc.mutation.Clubstatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clubappstatus.FieldClubstatus,
		})
		cs.Clubstatus = value
	}
	if nodes := csc.mutation.ClubapplicationIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return cs, _spec
}
