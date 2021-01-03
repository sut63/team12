// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/OMENX/app/ent/activities"
	"github.com/OMENX/app/ent/activitytype"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ActivityTypeCreate is the builder for creating a ActivityType entity.
type ActivityTypeCreate struct {
	config
	mutation *ActivityTypeMutation
	hooks    []Hook
}

// SetName sets the name field.
func (atc *ActivityTypeCreate) SetName(s string) *ActivityTypeCreate {
	atc.mutation.SetName(s)
	return atc
}

// AddActivityIDs adds the activities edge to Activities by ids.
func (atc *ActivityTypeCreate) AddActivityIDs(ids ...int) *ActivityTypeCreate {
	atc.mutation.AddActivityIDs(ids...)
	return atc
}

// AddActivities adds the activities edges to Activities.
func (atc *ActivityTypeCreate) AddActivities(a ...*Activities) *ActivityTypeCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atc.AddActivityIDs(ids...)
}

// Mutation returns the ActivityTypeMutation object of the builder.
func (atc *ActivityTypeCreate) Mutation() *ActivityTypeMutation {
	return atc.mutation
}

// Save creates the ActivityType in the database.
func (atc *ActivityTypeCreate) Save(ctx context.Context) (*ActivityType, error) {
	if _, ok := atc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := atc.mutation.Name(); ok {
		if err := activitytype.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	var (
		err  error
		node *ActivityType
	)
	if len(atc.hooks) == 0 {
		node, err = atc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActivityTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			atc.mutation = mutation
			node, err = atc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(atc.hooks) - 1; i >= 0; i-- {
			mut = atc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (atc *ActivityTypeCreate) SaveX(ctx context.Context) *ActivityType {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (atc *ActivityTypeCreate) sqlSave(ctx context.Context) (*ActivityType, error) {
	at, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	at.ID = int(id)
	return at, nil
}

func (atc *ActivityTypeCreate) createSpec() (*ActivityType, *sqlgraph.CreateSpec) {
	var (
		at    = &ActivityType{config: atc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: activitytype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: activitytype.FieldID,
			},
		}
	)
	if value, ok := atc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: activitytype.FieldName,
		})
		at.Name = value
	}
	if nodes := atc.mutation.ActivitiesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return at, _spec
}
