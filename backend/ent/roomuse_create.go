// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/OMENX/app/ent/purpose"
	"github.com/OMENX/app/ent/room"
	"github.com/OMENX/app/ent/roomuse"
	"github.com/OMENX/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RoomuseCreate is the builder for creating a Roomuse entity.
type RoomuseCreate struct {
	config
	mutation *RoomuseMutation
	hooks    []Hook
}

// SetAge sets the age field.
func (rc *RoomuseCreate) SetAge(i int) *RoomuseCreate {
	rc.mutation.SetAge(i)
	return rc
}

// SetNote sets the note field.
func (rc *RoomuseCreate) SetNote(s string) *RoomuseCreate {
	rc.mutation.SetNote(s)
	return rc
}

// SetContact sets the contact field.
func (rc *RoomuseCreate) SetContact(s string) *RoomuseCreate {
	rc.mutation.SetContact(s)
	return rc
}

// SetInTime sets the in_time field.
func (rc *RoomuseCreate) SetInTime(t time.Time) *RoomuseCreate {
	rc.mutation.SetInTime(t)
	return rc
}

// SetOutTime sets the out_time field.
func (rc *RoomuseCreate) SetOutTime(t time.Time) *RoomuseCreate {
	rc.mutation.SetOutTime(t)
	return rc
}

// SetRoomsID sets the rooms edge to Room by id.
func (rc *RoomuseCreate) SetRoomsID(id int) *RoomuseCreate {
	rc.mutation.SetRoomsID(id)
	return rc
}

// SetNillableRoomsID sets the rooms edge to Room by id if the given value is not nil.
func (rc *RoomuseCreate) SetNillableRoomsID(id *int) *RoomuseCreate {
	if id != nil {
		rc = rc.SetRoomsID(*id)
	}
	return rc
}

// SetRooms sets the rooms edge to Room.
func (rc *RoomuseCreate) SetRooms(r *Room) *RoomuseCreate {
	return rc.SetRoomsID(r.ID)
}

// SetPurposesID sets the purposes edge to Purpose by id.
func (rc *RoomuseCreate) SetPurposesID(id int) *RoomuseCreate {
	rc.mutation.SetPurposesID(id)
	return rc
}

// SetNillablePurposesID sets the purposes edge to Purpose by id if the given value is not nil.
func (rc *RoomuseCreate) SetNillablePurposesID(id *int) *RoomuseCreate {
	if id != nil {
		rc = rc.SetPurposesID(*id)
	}
	return rc
}

// SetPurposes sets the purposes edge to Purpose.
func (rc *RoomuseCreate) SetPurposes(p *Purpose) *RoomuseCreate {
	return rc.SetPurposesID(p.ID)
}

// SetUsersID sets the users edge to User by id.
func (rc *RoomuseCreate) SetUsersID(id int) *RoomuseCreate {
	rc.mutation.SetUsersID(id)
	return rc
}

// SetNillableUsersID sets the users edge to User by id if the given value is not nil.
func (rc *RoomuseCreate) SetNillableUsersID(id *int) *RoomuseCreate {
	if id != nil {
		rc = rc.SetUsersID(*id)
	}
	return rc
}

// SetUsers sets the users edge to User.
func (rc *RoomuseCreate) SetUsers(u *User) *RoomuseCreate {
	return rc.SetUsersID(u.ID)
}

// Mutation returns the RoomuseMutation object of the builder.
func (rc *RoomuseCreate) Mutation() *RoomuseMutation {
	return rc.mutation
}

// Save creates the Roomuse in the database.
func (rc *RoomuseCreate) Save(ctx context.Context) (*Roomuse, error) {
	if _, ok := rc.mutation.Age(); !ok {
		return nil, &ValidationError{Name: "age", err: errors.New("ent: missing required field \"age\"")}
	}
	if v, ok := rc.mutation.Age(); ok {
		if err := roomuse.AgeValidator(v); err != nil {
			return nil, &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	if _, ok := rc.mutation.Note(); !ok {
		return nil, &ValidationError{Name: "note", err: errors.New("ent: missing required field \"note\"")}
	}
	if v, ok := rc.mutation.Note(); ok {
		if err := roomuse.NoteValidator(v); err != nil {
			return nil, &ValidationError{Name: "note", err: fmt.Errorf("ent: validator failed for field \"note\": %w", err)}
		}
	}
	if _, ok := rc.mutation.Contact(); !ok {
		return nil, &ValidationError{Name: "contact", err: errors.New("ent: missing required field \"contact\"")}
	}
	if v, ok := rc.mutation.Contact(); ok {
		if err := roomuse.ContactValidator(v); err != nil {
			return nil, &ValidationError{Name: "contact", err: fmt.Errorf("ent: validator failed for field \"contact\": %w", err)}
		}
	}
	if _, ok := rc.mutation.InTime(); !ok {
		return nil, &ValidationError{Name: "in_time", err: errors.New("ent: missing required field \"in_time\"")}
	}
	if _, ok := rc.mutation.OutTime(); !ok {
		return nil, &ValidationError{Name: "out_time", err: errors.New("ent: missing required field \"out_time\"")}
	}
	var (
		err  error
		node *Roomuse
	)
	if len(rc.hooks) == 0 {
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomuseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rc.mutation = mutation
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoomuseCreate) SaveX(ctx context.Context) *Roomuse {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *RoomuseCreate) sqlSave(ctx context.Context) (*Roomuse, error) {
	r, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	r.ID = int(id)
	return r, nil
}

func (rc *RoomuseCreate) createSpec() (*Roomuse, *sqlgraph.CreateSpec) {
	var (
		r     = &Roomuse{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: roomuse.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roomuse.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roomuse.FieldAge,
		})
		r.Age = value
	}
	if value, ok := rc.mutation.Note(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roomuse.FieldNote,
		})
		r.Note = value
	}
	if value, ok := rc.mutation.Contact(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roomuse.FieldContact,
		})
		r.Contact = value
	}
	if value, ok := rc.mutation.InTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: roomuse.FieldInTime,
		})
		r.InTime = value
	}
	if value, ok := rc.mutation.OutTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: roomuse.FieldOutTime,
		})
		r.OutTime = value
	}
	if nodes := rc.mutation.RoomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roomuse.RoomsTable,
			Columns: []string{roomuse.RoomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.PurposesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roomuse.PurposesTable,
			Columns: []string{roomuse.PurposesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: purpose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roomuse.UsersTable,
			Columns: []string{roomuse.UsersColumn},
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
	return r, _spec
}
