// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/OMENX/app/ent/club"
	"github.com/OMENX/app/ent/discipline"
	"github.com/OMENX/app/ent/gender"
	"github.com/OMENX/app/ent/user"
	"github.com/OMENX/app/ent/userstatus"
	"github.com/OMENX/app/ent/usertype"
	"github.com/OMENX/app/ent/year"
	"github.com/facebookincubator/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges         UserEdges `json:"edges"`
	ClubID        *int
	discipline_id *int
	gender_id     *int
	userstatus_id *int
	UserTypeID    *int
	year_id       *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Usertype holds the value of the usertype edge.
	Usertype *Usertype
	// Clubuser holds the value of the clubuser edge.
	Clubuser *Club
	// Gender holds the value of the gender edge.
	Gender *Gender
	// Userstatus holds the value of the userstatus edge.
	Userstatus *UserStatus
	// Discipline holds the value of the discipline edge.
	Discipline *Discipline
	// Year holds the value of the year edge.
	Year *Year
	// Club holds the value of the club edge.
	Club []*Club
	// Clubapplication holds the value of the clubapplication edge.
	Clubapplication []*Clubapplication
	// UserToComplaint holds the value of the UserToComplaint edge.
	UserToComplaint []*Complaint
	// Roomuse holds the value of the Roomuse edge.
	Roomuse []*Roomuse
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [10]bool
}

// UsertypeOrErr returns the Usertype value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) UsertypeOrErr() (*Usertype, error) {
	if e.loadedTypes[0] {
		if e.Usertype == nil {
			// The edge usertype was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: usertype.Label}
		}
		return e.Usertype, nil
	}
	return nil, &NotLoadedError{edge: "usertype"}
}

// ClubuserOrErr returns the Clubuser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ClubuserOrErr() (*Club, error) {
	if e.loadedTypes[1] {
		if e.Clubuser == nil {
			// The edge clubuser was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: club.Label}
		}
		return e.Clubuser, nil
	}
	return nil, &NotLoadedError{edge: "clubuser"}
}

// GenderOrErr returns the Gender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) GenderOrErr() (*Gender, error) {
	if e.loadedTypes[2] {
		if e.Gender == nil {
			// The edge gender was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: gender.Label}
		}
		return e.Gender, nil
	}
	return nil, &NotLoadedError{edge: "gender"}
}

// UserstatusOrErr returns the Userstatus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) UserstatusOrErr() (*UserStatus, error) {
	if e.loadedTypes[3] {
		if e.Userstatus == nil {
			// The edge userstatus was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: userstatus.Label}
		}
		return e.Userstatus, nil
	}
	return nil, &NotLoadedError{edge: "userstatus"}
}

// DisciplineOrErr returns the Discipline value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) DisciplineOrErr() (*Discipline, error) {
	if e.loadedTypes[4] {
		if e.Discipline == nil {
			// The edge discipline was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: discipline.Label}
		}
		return e.Discipline, nil
	}
	return nil, &NotLoadedError{edge: "discipline"}
}

// YearOrErr returns the Year value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) YearOrErr() (*Year, error) {
	if e.loadedTypes[5] {
		if e.Year == nil {
			// The edge year was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: year.Label}
		}
		return e.Year, nil
	}
	return nil, &NotLoadedError{edge: "year"}
}

// ClubOrErr returns the Club value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ClubOrErr() ([]*Club, error) {
	if e.loadedTypes[6] {
		return e.Club, nil
	}
	return nil, &NotLoadedError{edge: "club"}
}

// ClubapplicationOrErr returns the Clubapplication value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ClubapplicationOrErr() ([]*Clubapplication, error) {
	if e.loadedTypes[7] {
		return e.Clubapplication, nil
	}
	return nil, &NotLoadedError{edge: "clubapplication"}
}

// UserToComplaintOrErr returns the UserToComplaint value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserToComplaintOrErr() ([]*Complaint, error) {
	if e.loadedTypes[8] {
		return e.UserToComplaint, nil
	}
	return nil, &NotLoadedError{edge: "UserToComplaint"}
}

// RoomuseOrErr returns the Roomuse value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RoomuseOrErr() ([]*Roomuse, error) {
	if e.loadedTypes[9] {
		return e.Roomuse, nil
	}
	return nil, &NotLoadedError{edge: "Roomuse"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // email
		&sql.NullString{}, // password
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*User) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // ClubID
		&sql.NullInt64{}, // discipline_id
		&sql.NullInt64{}, // gender_id
		&sql.NullInt64{}, // userstatus_id
		&sql.NullInt64{}, // UserTypeID
		&sql.NullInt64{}, // year_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		u.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[1])
	} else if value.Valid {
		u.Email = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[2])
	} else if value.Valid {
		u.Password = value.String
	}
	values = values[3:]
	if len(values) == len(user.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field ClubID", value)
		} else if value.Valid {
			u.ClubID = new(int)
			*u.ClubID = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field discipline_id", value)
		} else if value.Valid {
			u.discipline_id = new(int)
			*u.discipline_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field gender_id", value)
		} else if value.Valid {
			u.gender_id = new(int)
			*u.gender_id = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field userstatus_id", value)
		} else if value.Valid {
			u.userstatus_id = new(int)
			*u.userstatus_id = int(value.Int64)
		}
		if value, ok := values[4].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field UserTypeID", value)
		} else if value.Valid {
			u.UserTypeID = new(int)
			*u.UserTypeID = int(value.Int64)
		}
		if value, ok := values[5].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field year_id", value)
		} else if value.Valid {
			u.year_id = new(int)
			*u.year_id = int(value.Int64)
		}
	}
	return nil
}

// QueryUsertype queries the usertype edge of the User.
func (u *User) QueryUsertype() *UsertypeQuery {
	return (&UserClient{config: u.config}).QueryUsertype(u)
}

// QueryClubuser queries the clubuser edge of the User.
func (u *User) QueryClubuser() *ClubQuery {
	return (&UserClient{config: u.config}).QueryClubuser(u)
}

// QueryGender queries the gender edge of the User.
func (u *User) QueryGender() *GenderQuery {
	return (&UserClient{config: u.config}).QueryGender(u)
}

// QueryUserstatus queries the userstatus edge of the User.
func (u *User) QueryUserstatus() *UserStatusQuery {
	return (&UserClient{config: u.config}).QueryUserstatus(u)
}

// QueryDiscipline queries the discipline edge of the User.
func (u *User) QueryDiscipline() *DisciplineQuery {
	return (&UserClient{config: u.config}).QueryDiscipline(u)
}

// QueryYear queries the year edge of the User.
func (u *User) QueryYear() *YearQuery {
	return (&UserClient{config: u.config}).QueryYear(u)
}

// QueryClub queries the club edge of the User.
func (u *User) QueryClub() *ClubQuery {
	return (&UserClient{config: u.config}).QueryClub(u)
}

// QueryClubapplication queries the clubapplication edge of the User.
func (u *User) QueryClubapplication() *ClubapplicationQuery {
	return (&UserClient{config: u.config}).QueryClubapplication(u)
}

// QueryUserToComplaint queries the UserToComplaint edge of the User.
func (u *User) QueryUserToComplaint() *ComplaintQuery {
	return (&UserClient{config: u.config}).QueryUserToComplaint(u)
}

// QueryRoomuse queries the Roomuse edge of the User.
func (u *User) QueryRoomuse() *RoomuseQuery {
	return (&UserClient{config: u.config}).QueryRoomuse(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", password=")
	builder.WriteString(u.Password)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
