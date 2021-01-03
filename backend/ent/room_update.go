// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/OMENX/app/ent/predicate"
	"github.com/OMENX/app/ent/room"
	"github.com/OMENX/app/ent/roomuse"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RoomUpdate is the builder for updating Room entities.
type RoomUpdate struct {
	config
	hooks      []Hook
	mutation   *RoomMutation
	predicates []predicate.Room
}

// Where adds a new predicate for the builder.
func (ru *RoomUpdate) Where(ps ...predicate.Room) *RoomUpdate {
	ru.predicates = append(ru.predicates, ps...)
	return ru
}

// SetRoomName sets the room_name field.
func (ru *RoomUpdate) SetRoomName(s string) *RoomUpdate {
	ru.mutation.SetRoomName(s)
	return ru
}

// SetRoomLocation sets the room_location field.
func (ru *RoomUpdate) SetRoomLocation(s string) *RoomUpdate {
	ru.mutation.SetRoomLocation(s)
	return ru
}

// SetRoomFloor sets the room_floor field.
func (ru *RoomUpdate) SetRoomFloor(s string) *RoomUpdate {
	ru.mutation.SetRoomFloor(s)
	return ru
}

// SetMaxContain sets the max_contain field.
func (ru *RoomUpdate) SetMaxContain(i int) *RoomUpdate {
	ru.mutation.ResetMaxContain()
	ru.mutation.SetMaxContain(i)
	return ru
}

// AddMaxContain adds i to max_contain.
func (ru *RoomUpdate) AddMaxContain(i int) *RoomUpdate {
	ru.mutation.AddMaxContain(i)
	return ru
}

// AddRoomuseIDs adds the roomuses edge to Roomuse by ids.
func (ru *RoomUpdate) AddRoomuseIDs(ids ...int) *RoomUpdate {
	ru.mutation.AddRoomuseIDs(ids...)
	return ru
}

// AddRoomuses adds the roomuses edges to Roomuse.
func (ru *RoomUpdate) AddRoomuses(r ...*Roomuse) *RoomUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddRoomuseIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (ru *RoomUpdate) Mutation() *RoomMutation {
	return ru.mutation
}

// RemoveRoomuseIDs removes the roomuses edge to Roomuse by ids.
func (ru *RoomUpdate) RemoveRoomuseIDs(ids ...int) *RoomUpdate {
	ru.mutation.RemoveRoomuseIDs(ids...)
	return ru
}

// RemoveRoomuses removes roomuses edges to Roomuse.
func (ru *RoomUpdate) RemoveRoomuses(r ...*Roomuse) *RoomUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveRoomuseIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *RoomUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := ru.mutation.RoomName(); ok {
		if err := room.RoomNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "room_name", err: fmt.Errorf("ent: validator failed for field \"room_name\": %w", err)}
		}
	}
	if v, ok := ru.mutation.RoomLocation(); ok {
		if err := room.RoomLocationValidator(v); err != nil {
			return 0, &ValidationError{Name: "room_location", err: fmt.Errorf("ent: validator failed for field \"room_location\": %w", err)}
		}
	}
	if v, ok := ru.mutation.RoomFloor(); ok {
		if err := room.RoomFloorValidator(v); err != nil {
			return 0, &ValidationError{Name: "room_floor", err: fmt.Errorf("ent: validator failed for field \"room_floor\": %w", err)}
		}
	}
	if v, ok := ru.mutation.MaxContain(); ok {
		if err := room.MaxContainValidator(v); err != nil {
			return 0, &ValidationError{Name: "max_contain", err: fmt.Errorf("ent: validator failed for field \"max_contain\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoomUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoomUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoomUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RoomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   room.Table,
			Columns: room.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: room.FieldID,
			},
		},
	}
	if ps := ru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.RoomName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: room.FieldRoomName,
		})
	}
	if value, ok := ru.mutation.RoomLocation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: room.FieldRoomLocation,
		})
	}
	if value, ok := ru.mutation.RoomFloor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: room.FieldRoomFloor,
		})
	}
	if value, ok := ru.mutation.MaxContain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: room.FieldMaxContain,
		})
	}
	if value, ok := ru.mutation.AddedMaxContain(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: room.FieldMaxContain,
		})
	}
	if nodes := ru.mutation.RemovedRoomusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.RoomusesTable,
			Columns: []string{room.RoomusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomuse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RoomusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.RoomusesTable,
			Columns: []string{room.RoomusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomuse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{room.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RoomUpdateOne is the builder for updating a single Room entity.
type RoomUpdateOne struct {
	config
	hooks    []Hook
	mutation *RoomMutation
}

// SetRoomName sets the room_name field.
func (ruo *RoomUpdateOne) SetRoomName(s string) *RoomUpdateOne {
	ruo.mutation.SetRoomName(s)
	return ruo
}

// SetRoomLocation sets the room_location field.
func (ruo *RoomUpdateOne) SetRoomLocation(s string) *RoomUpdateOne {
	ruo.mutation.SetRoomLocation(s)
	return ruo
}

// SetRoomFloor sets the room_floor field.
func (ruo *RoomUpdateOne) SetRoomFloor(s string) *RoomUpdateOne {
	ruo.mutation.SetRoomFloor(s)
	return ruo
}

// SetMaxContain sets the max_contain field.
func (ruo *RoomUpdateOne) SetMaxContain(i int) *RoomUpdateOne {
	ruo.mutation.ResetMaxContain()
	ruo.mutation.SetMaxContain(i)
	return ruo
}

// AddMaxContain adds i to max_contain.
func (ruo *RoomUpdateOne) AddMaxContain(i int) *RoomUpdateOne {
	ruo.mutation.AddMaxContain(i)
	return ruo
}

// AddRoomuseIDs adds the roomuses edge to Roomuse by ids.
func (ruo *RoomUpdateOne) AddRoomuseIDs(ids ...int) *RoomUpdateOne {
	ruo.mutation.AddRoomuseIDs(ids...)
	return ruo
}

// AddRoomuses adds the roomuses edges to Roomuse.
func (ruo *RoomUpdateOne) AddRoomuses(r ...*Roomuse) *RoomUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddRoomuseIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (ruo *RoomUpdateOne) Mutation() *RoomMutation {
	return ruo.mutation
}

// RemoveRoomuseIDs removes the roomuses edge to Roomuse by ids.
func (ruo *RoomUpdateOne) RemoveRoomuseIDs(ids ...int) *RoomUpdateOne {
	ruo.mutation.RemoveRoomuseIDs(ids...)
	return ruo
}

// RemoveRoomuses removes roomuses edges to Roomuse.
func (ruo *RoomUpdateOne) RemoveRoomuses(r ...*Roomuse) *RoomUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveRoomuseIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (ruo *RoomUpdateOne) Save(ctx context.Context) (*Room, error) {
	if v, ok := ruo.mutation.RoomName(); ok {
		if err := room.RoomNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "room_name", err: fmt.Errorf("ent: validator failed for field \"room_name\": %w", err)}
		}
	}
	if v, ok := ruo.mutation.RoomLocation(); ok {
		if err := room.RoomLocationValidator(v); err != nil {
			return nil, &ValidationError{Name: "room_location", err: fmt.Errorf("ent: validator failed for field \"room_location\": %w", err)}
		}
	}
	if v, ok := ruo.mutation.RoomFloor(); ok {
		if err := room.RoomFloorValidator(v); err != nil {
			return nil, &ValidationError{Name: "room_floor", err: fmt.Errorf("ent: validator failed for field \"room_floor\": %w", err)}
		}
	}
	if v, ok := ruo.mutation.MaxContain(); ok {
		if err := room.MaxContainValidator(v); err != nil {
			return nil, &ValidationError{Name: "max_contain", err: fmt.Errorf("ent: validator failed for field \"max_contain\": %w", err)}
		}
	}

	var (
		err  error
		node *Room
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoomUpdateOne) SaveX(ctx context.Context) *Room {
	r, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// Exec executes the query on the entity.
func (ruo *RoomUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoomUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RoomUpdateOne) sqlSave(ctx context.Context) (r *Room, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   room.Table,
			Columns: room.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: room.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Room.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ruo.mutation.RoomName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: room.FieldRoomName,
		})
	}
	if value, ok := ruo.mutation.RoomLocation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: room.FieldRoomLocation,
		})
	}
	if value, ok := ruo.mutation.RoomFloor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: room.FieldRoomFloor,
		})
	}
	if value, ok := ruo.mutation.MaxContain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: room.FieldMaxContain,
		})
	}
	if value, ok := ruo.mutation.AddedMaxContain(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: room.FieldMaxContain,
		})
	}
	if nodes := ruo.mutation.RemovedRoomusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.RoomusesTable,
			Columns: []string{room.RoomusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomuse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RoomusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.RoomusesTable,
			Columns: []string{room.RoomusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomuse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	r = &Room{config: ruo.config}
	_spec.Assign = r.assignValues
	_spec.ScanValues = r.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{room.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return r, nil
}
