// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/OMENX/app/ent/usertype"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Usertype is the model entity for the Usertype schema.
type Usertype struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UsertypeQuery when eager-loading is set.
	Edges UsertypeEdges `json:"edges"`
}

// UsertypeEdges holds the relations/edges for other nodes in the graph.
type UsertypeEdges struct {
	// User holds the value of the user edge.
	User []*User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e UsertypeEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Usertype) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Usertype fields.
func (u *Usertype) assignValues(values ...interface{}) error {
	if m, n := len(values), len(usertype.Columns); m < n {
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
	return nil
}

// QueryUser queries the user edge of the Usertype.
func (u *Usertype) QueryUser() *UserQuery {
	return (&UsertypeClient{config: u.config}).QueryUser(u)
}

// Update returns a builder for updating this Usertype.
// Note that, you need to call Usertype.Unwrap() before calling this method, if this Usertype
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Usertype) Update() *UsertypeUpdateOne {
	return (&UsertypeClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *Usertype) Unwrap() *Usertype {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Usertype is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Usertype) String() string {
	var builder strings.Builder
	builder.WriteString("Usertype(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Usertypes is a parsable slice of Usertype.
type Usertypes []*Usertype

func (u Usertypes) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}