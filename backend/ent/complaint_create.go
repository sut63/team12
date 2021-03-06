// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/OMENX/app/ent/club"
	"github.com/OMENX/app/ent/complaint"
	"github.com/OMENX/app/ent/complainttype"
	"github.com/OMENX/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ComplaintCreate is the builder for creating a Complaint entity.
type ComplaintCreate struct {
	config
	mutation *ComplaintMutation
	hooks    []Hook
}

// SetName sets the name field.
func (cc *ComplaintCreate) SetName(s string) *ComplaintCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetPhonenumber sets the phonenumber field.
func (cc *ComplaintCreate) SetPhonenumber(s string) *ComplaintCreate {
	cc.mutation.SetPhonenumber(s)
	return cc
}

// SetInfo sets the info field.
func (cc *ComplaintCreate) SetInfo(s string) *ComplaintCreate {
	cc.mutation.SetInfo(s)
	return cc
}

// SetDate sets the date field.
func (cc *ComplaintCreate) SetDate(t time.Time) *ComplaintCreate {
	cc.mutation.SetDate(t)
	return cc
}

// SetComplaintToUserID sets the ComplaintToUser edge to User by id.
func (cc *ComplaintCreate) SetComplaintToUserID(id int) *ComplaintCreate {
	cc.mutation.SetComplaintToUserID(id)
	return cc
}

// SetNillableComplaintToUserID sets the ComplaintToUser edge to User by id if the given value is not nil.
func (cc *ComplaintCreate) SetNillableComplaintToUserID(id *int) *ComplaintCreate {
	if id != nil {
		cc = cc.SetComplaintToUserID(*id)
	}
	return cc
}

// SetComplaintToUser sets the ComplaintToUser edge to User.
func (cc *ComplaintCreate) SetComplaintToUser(u *User) *ComplaintCreate {
	return cc.SetComplaintToUserID(u.ID)
}

// SetComplaintToComplaintTypeID sets the ComplaintToComplaintType edge to ComplaintType by id.
func (cc *ComplaintCreate) SetComplaintToComplaintTypeID(id int) *ComplaintCreate {
	cc.mutation.SetComplaintToComplaintTypeID(id)
	return cc
}

// SetNillableComplaintToComplaintTypeID sets the ComplaintToComplaintType edge to ComplaintType by id if the given value is not nil.
func (cc *ComplaintCreate) SetNillableComplaintToComplaintTypeID(id *int) *ComplaintCreate {
	if id != nil {
		cc = cc.SetComplaintToComplaintTypeID(*id)
	}
	return cc
}

// SetComplaintToComplaintType sets the ComplaintToComplaintType edge to ComplaintType.
func (cc *ComplaintCreate) SetComplaintToComplaintType(c *ComplaintType) *ComplaintCreate {
	return cc.SetComplaintToComplaintTypeID(c.ID)
}

// SetComplaintToClubID sets the ComplaintToClub edge to Club by id.
func (cc *ComplaintCreate) SetComplaintToClubID(id int) *ComplaintCreate {
	cc.mutation.SetComplaintToClubID(id)
	return cc
}

// SetNillableComplaintToClubID sets the ComplaintToClub edge to Club by id if the given value is not nil.
func (cc *ComplaintCreate) SetNillableComplaintToClubID(id *int) *ComplaintCreate {
	if id != nil {
		cc = cc.SetComplaintToClubID(*id)
	}
	return cc
}

// SetComplaintToClub sets the ComplaintToClub edge to Club.
func (cc *ComplaintCreate) SetComplaintToClub(c *Club) *ComplaintCreate {
	return cc.SetComplaintToClubID(c.ID)
}

// Mutation returns the ComplaintMutation object of the builder.
func (cc *ComplaintCreate) Mutation() *ComplaintMutation {
	return cc.mutation
}

// Save creates the Complaint in the database.
func (cc *ComplaintCreate) Save(ctx context.Context) (*Complaint, error) {
	if _, ok := cc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := complaint.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := cc.mutation.Phonenumber(); !ok {
		return nil, &ValidationError{Name: "phonenumber", err: errors.New("ent: missing required field \"phonenumber\"")}
	}
	if v, ok := cc.mutation.Phonenumber(); ok {
		if err := complaint.PhonenumberValidator(v); err != nil {
			return nil, &ValidationError{Name: "phonenumber", err: fmt.Errorf("ent: validator failed for field \"phonenumber\": %w", err)}
		}
	}
	if _, ok := cc.mutation.Info(); !ok {
		return nil, &ValidationError{Name: "info", err: errors.New("ent: missing required field \"info\"")}
	}
	if v, ok := cc.mutation.Info(); ok {
		if err := complaint.InfoValidator(v); err != nil {
			return nil, &ValidationError{Name: "info", err: fmt.Errorf("ent: validator failed for field \"info\": %w", err)}
		}
	}
	if _, ok := cc.mutation.Date(); !ok {
		return nil, &ValidationError{Name: "date", err: errors.New("ent: missing required field \"date\"")}
	}
	var (
		err  error
		node *Complaint
	)
	if len(cc.hooks) == 0 {
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ComplaintMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ComplaintCreate) SaveX(ctx context.Context) *Complaint {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *ComplaintCreate) sqlSave(ctx context.Context) (*Complaint, error) {
	c, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	c.ID = int(id)
	return c, nil
}

func (cc *ComplaintCreate) createSpec() (*Complaint, *sqlgraph.CreateSpec) {
	var (
		c     = &Complaint{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: complaint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: complaint.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: complaint.FieldName,
		})
		c.Name = value
	}
	if value, ok := cc.mutation.Phonenumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: complaint.FieldPhonenumber,
		})
		c.Phonenumber = value
	}
	if value, ok := cc.mutation.Info(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: complaint.FieldInfo,
		})
		c.Info = value
	}
	if value, ok := cc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: complaint.FieldDate,
		})
		c.Date = value
	}
	if nodes := cc.mutation.ComplaintToUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   complaint.ComplaintToUserTable,
			Columns: []string{complaint.ComplaintToUserColumn},
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
	if nodes := cc.mutation.ComplaintToComplaintTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   complaint.ComplaintToComplaintTypeTable,
			Columns: []string{complaint.ComplaintToComplaintTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: complainttype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ComplaintToClubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   complaint.ComplaintToClubTable,
			Columns: []string{complaint.ComplaintToClubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return c, _spec
}
