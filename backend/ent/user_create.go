// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/OMENX/app/ent/club"
	"github.com/OMENX/app/ent/clubapplication"
	"github.com/OMENX/app/ent/complaint"
	"github.com/OMENX/app/ent/discipline"
	"github.com/OMENX/app/ent/gender"
	"github.com/OMENX/app/ent/roomuse"
	"github.com/OMENX/app/ent/user"
	"github.com/OMENX/app/ent/userstatus"
	"github.com/OMENX/app/ent/usertype"
	"github.com/OMENX/app/ent/year"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetName sets the name field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetAge sets the age field.
func (uc *UserCreate) SetAge(i int) *UserCreate {
	uc.mutation.SetAge(i)
	return uc
}

// SetEmail sets the email field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetPassword sets the password field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetUsertypeID sets the usertype edge to Usertype by id.
func (uc *UserCreate) SetUsertypeID(id int) *UserCreate {
	uc.mutation.SetUsertypeID(id)
	return uc
}

// SetNillableUsertypeID sets the usertype edge to Usertype by id if the given value is not nil.
func (uc *UserCreate) SetNillableUsertypeID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetUsertypeID(*id)
	}
	return uc
}

// SetUsertype sets the usertype edge to Usertype.
func (uc *UserCreate) SetUsertype(u *Usertype) *UserCreate {
	return uc.SetUsertypeID(u.ID)
}

// SetClubuserID sets the clubuser edge to Club by id.
func (uc *UserCreate) SetClubuserID(id int) *UserCreate {
	uc.mutation.SetClubuserID(id)
	return uc
}

// SetNillableClubuserID sets the clubuser edge to Club by id if the given value is not nil.
func (uc *UserCreate) SetNillableClubuserID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetClubuserID(*id)
	}
	return uc
}

// SetClubuser sets the clubuser edge to Club.
func (uc *UserCreate) SetClubuser(c *Club) *UserCreate {
	return uc.SetClubuserID(c.ID)
}

// SetGenderID sets the gender edge to Gender by id.
func (uc *UserCreate) SetGenderID(id int) *UserCreate {
	uc.mutation.SetGenderID(id)
	return uc
}

// SetNillableGenderID sets the gender edge to Gender by id if the given value is not nil.
func (uc *UserCreate) SetNillableGenderID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetGenderID(*id)
	}
	return uc
}

// SetGender sets the gender edge to Gender.
func (uc *UserCreate) SetGender(g *Gender) *UserCreate {
	return uc.SetGenderID(g.ID)
}

// SetUserstatusID sets the userstatus edge to UserStatus by id.
func (uc *UserCreate) SetUserstatusID(id int) *UserCreate {
	uc.mutation.SetUserstatusID(id)
	return uc
}

// SetNillableUserstatusID sets the userstatus edge to UserStatus by id if the given value is not nil.
func (uc *UserCreate) SetNillableUserstatusID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetUserstatusID(*id)
	}
	return uc
}

// SetUserstatus sets the userstatus edge to UserStatus.
func (uc *UserCreate) SetUserstatus(u *UserStatus) *UserCreate {
	return uc.SetUserstatusID(u.ID)
}

// SetDisciplineID sets the discipline edge to Discipline by id.
func (uc *UserCreate) SetDisciplineID(id int) *UserCreate {
	uc.mutation.SetDisciplineID(id)
	return uc
}

// SetNillableDisciplineID sets the discipline edge to Discipline by id if the given value is not nil.
func (uc *UserCreate) SetNillableDisciplineID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetDisciplineID(*id)
	}
	return uc
}

// SetDiscipline sets the discipline edge to Discipline.
func (uc *UserCreate) SetDiscipline(d *Discipline) *UserCreate {
	return uc.SetDisciplineID(d.ID)
}

// SetYearID sets the year edge to Year by id.
func (uc *UserCreate) SetYearID(id int) *UserCreate {
	uc.mutation.SetYearID(id)
	return uc
}

// SetNillableYearID sets the year edge to Year by id if the given value is not nil.
func (uc *UserCreate) SetNillableYearID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetYearID(*id)
	}
	return uc
}

// SetYear sets the year edge to Year.
func (uc *UserCreate) SetYear(y *Year) *UserCreate {
	return uc.SetYearID(y.ID)
}

// AddClubIDs adds the club edge to Club by ids.
func (uc *UserCreate) AddClubIDs(ids ...int) *UserCreate {
	uc.mutation.AddClubIDs(ids...)
	return uc
}

// AddClub adds the club edges to Club.
func (uc *UserCreate) AddClub(c ...*Club) *UserCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddClubIDs(ids...)
}

// AddClubapplicationIDs adds the clubapplication edge to Clubapplication by ids.
func (uc *UserCreate) AddClubapplicationIDs(ids ...int) *UserCreate {
	uc.mutation.AddClubapplicationIDs(ids...)
	return uc
}

// AddClubapplication adds the clubapplication edges to Clubapplication.
func (uc *UserCreate) AddClubapplication(c ...*Clubapplication) *UserCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddClubapplicationIDs(ids...)
}

// AddUserToComplaintIDs adds the UserToComplaint edge to Complaint by ids.
func (uc *UserCreate) AddUserToComplaintIDs(ids ...int) *UserCreate {
	uc.mutation.AddUserToComplaintIDs(ids...)
	return uc
}

// AddUserToComplaint adds the UserToComplaint edges to Complaint.
func (uc *UserCreate) AddUserToComplaint(c ...*Complaint) *UserCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddUserToComplaintIDs(ids...)
}

// AddRoomuseIDs adds the Roomuse edge to Roomuse by ids.
func (uc *UserCreate) AddRoomuseIDs(ids ...int) *UserCreate {
	uc.mutation.AddRoomuseIDs(ids...)
	return uc
}

// AddRoomuse adds the Roomuse edges to Roomuse.
func (uc *UserCreate) AddRoomuse(r ...*Roomuse) *UserCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddRoomuseIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if _, ok := uc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Age(); !ok {
		return nil, &ValidationError{Name: "age", err: errors.New("ent: missing required field \"age\"")}
	}
	if v, ok := uc.mutation.Age(); ok {
		if err := user.AgeValidator(v); err != nil {
			return nil, &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return nil, &ValidationError{Name: "email", err: errors.New("ent: missing required field \"email\"")}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return nil, &ValidationError{Name: "password", err: errors.New("ent: missing required field \"password\"")}
	}
	var (
		err  error
		node *User
	)
	if len(uc.hooks) == 0 {
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	u, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
		u.Name = value
	}
	if value, ok := uc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge,
		})
		u.Age = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
		u.Email = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
		u.Password = value
	}
	if nodes := uc.mutation.UsertypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.UsertypeTable,
			Columns: []string{user.UsertypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usertype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ClubuserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ClubuserTable,
			Columns: []string{user.ClubuserColumn},
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
	if nodes := uc.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.GenderTable,
			Columns: []string{user.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserstatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.UserstatusTable,
			Columns: []string{user.UserstatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.DisciplineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.DisciplineTable,
			Columns: []string{user.DisciplineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: discipline.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.YearIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.YearTable,
			Columns: []string{user.YearColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: year.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ClubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ClubTable,
			Columns: []string{user.ClubColumn},
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
	if nodes := uc.mutation.ClubapplicationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ClubapplicationTable,
			Columns: []string{user.ClubapplicationColumn},
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
	if nodes := uc.mutation.UserToComplaintIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserToComplaintTable,
			Columns: []string{user.UserToComplaintColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: complaint.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.RoomuseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RoomuseTable,
			Columns: []string{user.RoomuseColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return u, _spec
}
