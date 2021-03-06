// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/OMENX/app/ent/discipline"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Discipline is the model entity for the Discipline schema.
type Discipline struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Discipline holds the value of the "discipline" field.
	Discipline string `json:"discipline,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DisciplineQuery when eager-loading is set.
	Edges DisciplineEdges `json:"edges"`
}

// DisciplineEdges holds the relations/edges for other nodes in the graph.
type DisciplineEdges struct {
	// Users holds the value of the users edge.
	Users []*User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e DisciplineEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Discipline) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // discipline
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Discipline fields.
func (d *Discipline) assignValues(values ...interface{}) error {
	if m, n := len(values), len(discipline.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field discipline", values[0])
	} else if value.Valid {
		d.Discipline = value.String
	}
	return nil
}

// QueryUsers queries the users edge of the Discipline.
func (d *Discipline) QueryUsers() *UserQuery {
	return (&DisciplineClient{config: d.config}).QueryUsers(d)
}

// Update returns a builder for updating this Discipline.
// Note that, you need to call Discipline.Unwrap() before calling this method, if this Discipline
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Discipline) Update() *DisciplineUpdateOne {
	return (&DisciplineClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Discipline) Unwrap() *Discipline {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Discipline is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Discipline) String() string {
	var builder strings.Builder
	builder.WriteString("Discipline(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", discipline=")
	builder.WriteString(d.Discipline)
	builder.WriteByte(')')
	return builder.String()
}

// Disciplines is a parsable slice of Discipline.
type Disciplines []*Discipline

func (d Disciplines) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
