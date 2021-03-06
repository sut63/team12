// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/OMENX/app/ent/club"
	"github.com/OMENX/app/ent/clubapplication"
	"github.com/OMENX/app/ent/clubappstatus"
	"github.com/OMENX/app/ent/predicate"
	"github.com/OMENX/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ClubapplicationUpdate is the builder for updating Clubapplication entities.
type ClubapplicationUpdate struct {
	config
	hooks      []Hook
	mutation   *ClubapplicationMutation
	predicates []predicate.Clubapplication
}

// Where adds a new predicate for the builder.
func (cu *ClubapplicationUpdate) Where(ps ...predicate.Clubapplication) *ClubapplicationUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetContact sets the contact field.
func (cu *ClubapplicationUpdate) SetContact(s string) *ClubapplicationUpdate {
	cu.mutation.SetContact(s)
	return cu
}

// SetReason sets the reason field.
func (cu *ClubapplicationUpdate) SetReason(s string) *ClubapplicationUpdate {
	cu.mutation.SetReason(s)
	return cu
}

// SetAddeddatetime sets the addeddatetime field.
func (cu *ClubapplicationUpdate) SetAddeddatetime(t time.Time) *ClubapplicationUpdate {
	cu.mutation.SetAddeddatetime(t)
	return cu
}

// SetAddername sets the addername field.
func (cu *ClubapplicationUpdate) SetAddername(s string) *ClubapplicationUpdate {
	cu.mutation.SetAddername(s)
	return cu
}

// SetDiscipline sets the discipline field.
func (cu *ClubapplicationUpdate) SetDiscipline(s string) *ClubapplicationUpdate {
	cu.mutation.SetDiscipline(s)
	return cu
}

// SetGender sets the gender field.
func (cu *ClubapplicationUpdate) SetGender(s string) *ClubapplicationUpdate {
	cu.mutation.SetGender(s)
	return cu
}

// SetAge sets the age field.
func (cu *ClubapplicationUpdate) SetAge(i int) *ClubapplicationUpdate {
	cu.mutation.ResetAge()
	cu.mutation.SetAge(i)
	return cu
}

// AddAge adds i to age.
func (cu *ClubapplicationUpdate) AddAge(i int) *ClubapplicationUpdate {
	cu.mutation.AddAge(i)
	return cu
}

// SetYaer sets the yaer field.
func (cu *ClubapplicationUpdate) SetYaer(i int) *ClubapplicationUpdate {
	cu.mutation.ResetYaer()
	cu.mutation.SetYaer(i)
	return cu
}

// AddYaer adds i to yaer.
func (cu *ClubapplicationUpdate) AddYaer(i int) *ClubapplicationUpdate {
	cu.mutation.AddYaer(i)
	return cu
}

// SetOwnerID sets the owner edge to User by id.
func (cu *ClubapplicationUpdate) SetOwnerID(id int) *ClubapplicationUpdate {
	cu.mutation.SetOwnerID(id)
	return cu
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cu *ClubapplicationUpdate) SetNillableOwnerID(id *int) *ClubapplicationUpdate {
	if id != nil {
		cu = cu.SetOwnerID(*id)
	}
	return cu
}

// SetOwner sets the owner edge to User.
func (cu *ClubapplicationUpdate) SetOwner(u *User) *ClubapplicationUpdate {
	return cu.SetOwnerID(u.ID)
}

// SetClubappstatusID sets the clubappstatus edge to ClubappStatus by id.
func (cu *ClubapplicationUpdate) SetClubappstatusID(id int) *ClubapplicationUpdate {
	cu.mutation.SetClubappstatusID(id)
	return cu
}

// SetNillableClubappstatusID sets the clubappstatus edge to ClubappStatus by id if the given value is not nil.
func (cu *ClubapplicationUpdate) SetNillableClubappstatusID(id *int) *ClubapplicationUpdate {
	if id != nil {
		cu = cu.SetClubappstatusID(*id)
	}
	return cu
}

// SetClubappstatus sets the clubappstatus edge to ClubappStatus.
func (cu *ClubapplicationUpdate) SetClubappstatus(c *ClubappStatus) *ClubapplicationUpdate {
	return cu.SetClubappstatusID(c.ID)
}

// SetClubID sets the club edge to Club by id.
func (cu *ClubapplicationUpdate) SetClubID(id int) *ClubapplicationUpdate {
	cu.mutation.SetClubID(id)
	return cu
}

// SetNillableClubID sets the club edge to Club by id if the given value is not nil.
func (cu *ClubapplicationUpdate) SetNillableClubID(id *int) *ClubapplicationUpdate {
	if id != nil {
		cu = cu.SetClubID(*id)
	}
	return cu
}

// SetClub sets the club edge to Club.
func (cu *ClubapplicationUpdate) SetClub(c *Club) *ClubapplicationUpdate {
	return cu.SetClubID(c.ID)
}

// Mutation returns the ClubapplicationMutation object of the builder.
func (cu *ClubapplicationUpdate) Mutation() *ClubapplicationMutation {
	return cu.mutation
}

// ClearOwner clears the owner edge to User.
func (cu *ClubapplicationUpdate) ClearOwner() *ClubapplicationUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// ClearClubappstatus clears the clubappstatus edge to ClubappStatus.
func (cu *ClubapplicationUpdate) ClearClubappstatus() *ClubapplicationUpdate {
	cu.mutation.ClearClubappstatus()
	return cu
}

// ClearClub clears the club edge to Club.
func (cu *ClubapplicationUpdate) ClearClub() *ClubapplicationUpdate {
	cu.mutation.ClearClub()
	return cu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *ClubapplicationUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := cu.mutation.Contact(); ok {
		if err := clubapplication.ContactValidator(v); err != nil {
			return 0, &ValidationError{Name: "contact", err: fmt.Errorf("ent: validator failed for field \"contact\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Addername(); ok {
		if err := clubapplication.AddernameValidator(v); err != nil {
			return 0, &ValidationError{Name: "addername", err: fmt.Errorf("ent: validator failed for field \"addername\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Discipline(); ok {
		if err := clubapplication.DisciplineValidator(v); err != nil {
			return 0, &ValidationError{Name: "discipline", err: fmt.Errorf("ent: validator failed for field \"discipline\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Gender(); ok {
		if err := clubapplication.GenderValidator(v); err != nil {
			return 0, &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Age(); ok {
		if err := clubapplication.AgeValidator(v); err != nil {
			return 0, &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Yaer(); ok {
		if err := clubapplication.YaerValidator(v); err != nil {
			return 0, &ValidationError{Name: "yaer", err: fmt.Errorf("ent: validator failed for field \"yaer\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClubapplicationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ClubapplicationUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ClubapplicationUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ClubapplicationUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ClubapplicationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clubapplication.Table,
			Columns: clubapplication.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clubapplication.FieldID,
			},
		},
	}
	if ps := cu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Contact(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clubapplication.FieldContact,
		})
	}
	if value, ok := cu.mutation.Reason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clubapplication.FieldReason,
		})
	}
	if value, ok := cu.mutation.Addeddatetime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clubapplication.FieldAddeddatetime,
		})
	}
	if value, ok := cu.mutation.Addername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clubapplication.FieldAddername,
		})
	}
	if value, ok := cu.mutation.Discipline(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clubapplication.FieldDiscipline,
		})
	}
	if value, ok := cu.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clubapplication.FieldGender,
		})
	}
	if value, ok := cu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clubapplication.FieldAge,
		})
	}
	if value, ok := cu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clubapplication.FieldAge,
		})
	}
	if value, ok := cu.mutation.Yaer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clubapplication.FieldYaer,
		})
	}
	if value, ok := cu.mutation.AddedYaer(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clubapplication.FieldYaer,
		})
	}
	if cu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clubapplication.OwnerTable,
			Columns: []string{clubapplication.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clubapplication.OwnerTable,
			Columns: []string{clubapplication.OwnerColumn},
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
	if cu.mutation.ClubappstatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clubapplication.ClubappstatusTable,
			Columns: []string{clubapplication.ClubappstatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubappstatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ClubappstatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clubapplication.ClubappstatusTable,
			Columns: []string{clubapplication.ClubappstatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubappstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ClubCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clubapplication.ClubTable,
			Columns: []string{clubapplication.ClubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ClubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clubapplication.ClubTable,
			Columns: []string{clubapplication.ClubColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clubapplication.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ClubapplicationUpdateOne is the builder for updating a single Clubapplication entity.
type ClubapplicationUpdateOne struct {
	config
	hooks    []Hook
	mutation *ClubapplicationMutation
}

// SetContact sets the contact field.
func (cuo *ClubapplicationUpdateOne) SetContact(s string) *ClubapplicationUpdateOne {
	cuo.mutation.SetContact(s)
	return cuo
}

// SetReason sets the reason field.
func (cuo *ClubapplicationUpdateOne) SetReason(s string) *ClubapplicationUpdateOne {
	cuo.mutation.SetReason(s)
	return cuo
}

// SetAddeddatetime sets the addeddatetime field.
func (cuo *ClubapplicationUpdateOne) SetAddeddatetime(t time.Time) *ClubapplicationUpdateOne {
	cuo.mutation.SetAddeddatetime(t)
	return cuo
}

// SetAddername sets the addername field.
func (cuo *ClubapplicationUpdateOne) SetAddername(s string) *ClubapplicationUpdateOne {
	cuo.mutation.SetAddername(s)
	return cuo
}

// SetDiscipline sets the discipline field.
func (cuo *ClubapplicationUpdateOne) SetDiscipline(s string) *ClubapplicationUpdateOne {
	cuo.mutation.SetDiscipline(s)
	return cuo
}

// SetGender sets the gender field.
func (cuo *ClubapplicationUpdateOne) SetGender(s string) *ClubapplicationUpdateOne {
	cuo.mutation.SetGender(s)
	return cuo
}

// SetAge sets the age field.
func (cuo *ClubapplicationUpdateOne) SetAge(i int) *ClubapplicationUpdateOne {
	cuo.mutation.ResetAge()
	cuo.mutation.SetAge(i)
	return cuo
}

// AddAge adds i to age.
func (cuo *ClubapplicationUpdateOne) AddAge(i int) *ClubapplicationUpdateOne {
	cuo.mutation.AddAge(i)
	return cuo
}

// SetYaer sets the yaer field.
func (cuo *ClubapplicationUpdateOne) SetYaer(i int) *ClubapplicationUpdateOne {
	cuo.mutation.ResetYaer()
	cuo.mutation.SetYaer(i)
	return cuo
}

// AddYaer adds i to yaer.
func (cuo *ClubapplicationUpdateOne) AddYaer(i int) *ClubapplicationUpdateOne {
	cuo.mutation.AddYaer(i)
	return cuo
}

// SetOwnerID sets the owner edge to User by id.
func (cuo *ClubapplicationUpdateOne) SetOwnerID(id int) *ClubapplicationUpdateOne {
	cuo.mutation.SetOwnerID(id)
	return cuo
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cuo *ClubapplicationUpdateOne) SetNillableOwnerID(id *int) *ClubapplicationUpdateOne {
	if id != nil {
		cuo = cuo.SetOwnerID(*id)
	}
	return cuo
}

// SetOwner sets the owner edge to User.
func (cuo *ClubapplicationUpdateOne) SetOwner(u *User) *ClubapplicationUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// SetClubappstatusID sets the clubappstatus edge to ClubappStatus by id.
func (cuo *ClubapplicationUpdateOne) SetClubappstatusID(id int) *ClubapplicationUpdateOne {
	cuo.mutation.SetClubappstatusID(id)
	return cuo
}

// SetNillableClubappstatusID sets the clubappstatus edge to ClubappStatus by id if the given value is not nil.
func (cuo *ClubapplicationUpdateOne) SetNillableClubappstatusID(id *int) *ClubapplicationUpdateOne {
	if id != nil {
		cuo = cuo.SetClubappstatusID(*id)
	}
	return cuo
}

// SetClubappstatus sets the clubappstatus edge to ClubappStatus.
func (cuo *ClubapplicationUpdateOne) SetClubappstatus(c *ClubappStatus) *ClubapplicationUpdateOne {
	return cuo.SetClubappstatusID(c.ID)
}

// SetClubID sets the club edge to Club by id.
func (cuo *ClubapplicationUpdateOne) SetClubID(id int) *ClubapplicationUpdateOne {
	cuo.mutation.SetClubID(id)
	return cuo
}

// SetNillableClubID sets the club edge to Club by id if the given value is not nil.
func (cuo *ClubapplicationUpdateOne) SetNillableClubID(id *int) *ClubapplicationUpdateOne {
	if id != nil {
		cuo = cuo.SetClubID(*id)
	}
	return cuo
}

// SetClub sets the club edge to Club.
func (cuo *ClubapplicationUpdateOne) SetClub(c *Club) *ClubapplicationUpdateOne {
	return cuo.SetClubID(c.ID)
}

// Mutation returns the ClubapplicationMutation object of the builder.
func (cuo *ClubapplicationUpdateOne) Mutation() *ClubapplicationMutation {
	return cuo.mutation
}

// ClearOwner clears the owner edge to User.
func (cuo *ClubapplicationUpdateOne) ClearOwner() *ClubapplicationUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// ClearClubappstatus clears the clubappstatus edge to ClubappStatus.
func (cuo *ClubapplicationUpdateOne) ClearClubappstatus() *ClubapplicationUpdateOne {
	cuo.mutation.ClearClubappstatus()
	return cuo
}

// ClearClub clears the club edge to Club.
func (cuo *ClubapplicationUpdateOne) ClearClub() *ClubapplicationUpdateOne {
	cuo.mutation.ClearClub()
	return cuo
}

// Save executes the query and returns the updated entity.
func (cuo *ClubapplicationUpdateOne) Save(ctx context.Context) (*Clubapplication, error) {
	if v, ok := cuo.mutation.Contact(); ok {
		if err := clubapplication.ContactValidator(v); err != nil {
			return nil, &ValidationError{Name: "contact", err: fmt.Errorf("ent: validator failed for field \"contact\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Addername(); ok {
		if err := clubapplication.AddernameValidator(v); err != nil {
			return nil, &ValidationError{Name: "addername", err: fmt.Errorf("ent: validator failed for field \"addername\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Discipline(); ok {
		if err := clubapplication.DisciplineValidator(v); err != nil {
			return nil, &ValidationError{Name: "discipline", err: fmt.Errorf("ent: validator failed for field \"discipline\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Gender(); ok {
		if err := clubapplication.GenderValidator(v); err != nil {
			return nil, &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Age(); ok {
		if err := clubapplication.AgeValidator(v); err != nil {
			return nil, &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Yaer(); ok {
		if err := clubapplication.YaerValidator(v); err != nil {
			return nil, &ValidationError{Name: "yaer", err: fmt.Errorf("ent: validator failed for field \"yaer\": %w", err)}
		}
	}

	var (
		err  error
		node *Clubapplication
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClubapplicationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ClubapplicationUpdateOne) SaveX(ctx context.Context) *Clubapplication {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *ClubapplicationUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ClubapplicationUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ClubapplicationUpdateOne) sqlSave(ctx context.Context) (c *Clubapplication, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clubapplication.Table,
			Columns: clubapplication.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clubapplication.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Clubapplication.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.Contact(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clubapplication.FieldContact,
		})
	}
	if value, ok := cuo.mutation.Reason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clubapplication.FieldReason,
		})
	}
	if value, ok := cuo.mutation.Addeddatetime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clubapplication.FieldAddeddatetime,
		})
	}
	if value, ok := cuo.mutation.Addername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clubapplication.FieldAddername,
		})
	}
	if value, ok := cuo.mutation.Discipline(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clubapplication.FieldDiscipline,
		})
	}
	if value, ok := cuo.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clubapplication.FieldGender,
		})
	}
	if value, ok := cuo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clubapplication.FieldAge,
		})
	}
	if value, ok := cuo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clubapplication.FieldAge,
		})
	}
	if value, ok := cuo.mutation.Yaer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clubapplication.FieldYaer,
		})
	}
	if value, ok := cuo.mutation.AddedYaer(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clubapplication.FieldYaer,
		})
	}
	if cuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clubapplication.OwnerTable,
			Columns: []string{clubapplication.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clubapplication.OwnerTable,
			Columns: []string{clubapplication.OwnerColumn},
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
	if cuo.mutation.ClubappstatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clubapplication.ClubappstatusTable,
			Columns: []string{clubapplication.ClubappstatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubappstatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ClubappstatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clubapplication.ClubappstatusTable,
			Columns: []string{clubapplication.ClubappstatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubappstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ClubCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clubapplication.ClubTable,
			Columns: []string{clubapplication.ClubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ClubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clubapplication.ClubTable,
			Columns: []string{clubapplication.ClubColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	c = &Clubapplication{config: cuo.config}
	_spec.Assign = c.assignValues
	_spec.ScanValues = c.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clubapplication.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return c, nil
}
