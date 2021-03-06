// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/OMENX/app/ent/clubappstatus"
	"github.com/facebookincubator/ent/dialect/sql"
)

// ClubappStatus is the model entity for the ClubappStatus schema.
type ClubappStatus struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Clubstatus holds the value of the "clubstatus" field.
	Clubstatus string `json:"clubstatus,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClubappStatusQuery when eager-loading is set.
	Edges ClubappStatusEdges `json:"edges"`
}

// ClubappStatusEdges holds the relations/edges for other nodes in the graph.
type ClubappStatusEdges struct {
	// Clubapplication holds the value of the clubapplication edge.
	Clubapplication []*Clubapplication
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ClubapplicationOrErr returns the Clubapplication value or an error if the edge
// was not loaded in eager-loading.
func (e ClubappStatusEdges) ClubapplicationOrErr() ([]*Clubapplication, error) {
	if e.loadedTypes[0] {
		return e.Clubapplication, nil
	}
	return nil, &NotLoadedError{edge: "clubapplication"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ClubappStatus) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // clubstatus
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ClubappStatus fields.
func (cs *ClubappStatus) assignValues(values ...interface{}) error {
	if m, n := len(values), len(clubappstatus.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	cs.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field clubstatus", values[0])
	} else if value.Valid {
		cs.Clubstatus = value.String
	}
	return nil
}

// QueryClubapplication queries the clubapplication edge of the ClubappStatus.
func (cs *ClubappStatus) QueryClubapplication() *ClubapplicationQuery {
	return (&ClubappStatusClient{config: cs.config}).QueryClubapplication(cs)
}

// Update returns a builder for updating this ClubappStatus.
// Note that, you need to call ClubappStatus.Unwrap() before calling this method, if this ClubappStatus
// was returned from a transaction, and the transaction was committed or rolled back.
func (cs *ClubappStatus) Update() *ClubappStatusUpdateOne {
	return (&ClubappStatusClient{config: cs.config}).UpdateOne(cs)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (cs *ClubappStatus) Unwrap() *ClubappStatus {
	tx, ok := cs.config.driver.(*txDriver)
	if !ok {
		panic("ent: ClubappStatus is not a transactional entity")
	}
	cs.config.driver = tx.drv
	return cs
}

// String implements the fmt.Stringer.
func (cs *ClubappStatus) String() string {
	var builder strings.Builder
	builder.WriteString("ClubappStatus(")
	builder.WriteString(fmt.Sprintf("id=%v", cs.ID))
	builder.WriteString(", clubstatus=")
	builder.WriteString(cs.Clubstatus)
	builder.WriteByte(')')
	return builder.String()
}

// ClubappStatusSlice is a parsable slice of ClubappStatus.
type ClubappStatusSlice []*ClubappStatus

func (cs ClubappStatusSlice) config(cfg config) {
	for _i := range cs {
		cs[_i].config = cfg
	}
}
