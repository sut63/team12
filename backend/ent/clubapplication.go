// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/OMENX/app/ent/club"
	"github.com/OMENX/app/ent/clubapplication"
	"github.com/OMENX/app/ent/clubappstatus"
	"github.com/OMENX/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Clubapplication is the model entity for the Clubapplication schema.
type Clubapplication struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Contact holds the value of the "contact" field.
	Contact string `json:"contact,omitempty"`
	// Reason holds the value of the "reason" field.
	Reason string `json:"reason,omitempty"`
	// Addeddatetime holds the value of the "addeddatetime" field.
	Addeddatetime time.Time `json:"addeddatetime,omitempty"`
	// Addername holds the value of the "addername" field.
	Addername string `json:"addername,omitempty"`
	// Discipline holds the value of the "discipline" field.
	Discipline string `json:"discipline,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender string `json:"gender,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Yaer holds the value of the "yaer" field.
	Yaer int `json:"yaer,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClubapplicationQuery when eager-loading is set.
	Edges        ClubapplicationEdges `json:"edges"`
	ClubID       *int
	clubstatusID *int
	UserID       *int
}

// ClubapplicationEdges holds the relations/edges for other nodes in the graph.
type ClubapplicationEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User
	// Clubappstatus holds the value of the clubappstatus edge.
	Clubappstatus *ClubappStatus
	// Club holds the value of the club edge.
	Club *Club
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClubapplicationEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// ClubappstatusOrErr returns the Clubappstatus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClubapplicationEdges) ClubappstatusOrErr() (*ClubappStatus, error) {
	if e.loadedTypes[1] {
		if e.Clubappstatus == nil {
			// The edge clubappstatus was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: clubappstatus.Label}
		}
		return e.Clubappstatus, nil
	}
	return nil, &NotLoadedError{edge: "clubappstatus"}
}

// ClubOrErr returns the Club value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClubapplicationEdges) ClubOrErr() (*Club, error) {
	if e.loadedTypes[2] {
		if e.Club == nil {
			// The edge club was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: club.Label}
		}
		return e.Club, nil
	}
	return nil, &NotLoadedError{edge: "club"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Clubapplication) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // contact
		&sql.NullString{}, // reason
		&sql.NullTime{},   // addeddatetime
		&sql.NullString{}, // addername
		&sql.NullString{}, // discipline
		&sql.NullString{}, // gender
		&sql.NullInt64{},  // age
		&sql.NullInt64{},  // yaer
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Clubapplication) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // ClubID
		&sql.NullInt64{}, // clubstatusID
		&sql.NullInt64{}, // UserID
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Clubapplication fields.
func (c *Clubapplication) assignValues(values ...interface{}) error {
	if m, n := len(values), len(clubapplication.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field contact", values[0])
	} else if value.Valid {
		c.Contact = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field reason", values[1])
	} else if value.Valid {
		c.Reason = value.String
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field addeddatetime", values[2])
	} else if value.Valid {
		c.Addeddatetime = value.Time
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field addername", values[3])
	} else if value.Valid {
		c.Addername = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field discipline", values[4])
	} else if value.Valid {
		c.Discipline = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field gender", values[5])
	} else if value.Valid {
		c.Gender = value.String
	}
	if value, ok := values[6].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field age", values[6])
	} else if value.Valid {
		c.Age = int(value.Int64)
	}
	if value, ok := values[7].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field yaer", values[7])
	} else if value.Valid {
		c.Yaer = int(value.Int64)
	}
	values = values[8:]
	if len(values) == len(clubapplication.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field ClubID", value)
		} else if value.Valid {
			c.ClubID = new(int)
			*c.ClubID = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field clubstatusID", value)
		} else if value.Valid {
			c.clubstatusID = new(int)
			*c.clubstatusID = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field UserID", value)
		} else if value.Valid {
			c.UserID = new(int)
			*c.UserID = int(value.Int64)
		}
	}
	return nil
}

// QueryOwner queries the owner edge of the Clubapplication.
func (c *Clubapplication) QueryOwner() *UserQuery {
	return (&ClubapplicationClient{config: c.config}).QueryOwner(c)
}

// QueryClubappstatus queries the clubappstatus edge of the Clubapplication.
func (c *Clubapplication) QueryClubappstatus() *ClubappStatusQuery {
	return (&ClubapplicationClient{config: c.config}).QueryClubappstatus(c)
}

// QueryClub queries the club edge of the Clubapplication.
func (c *Clubapplication) QueryClub() *ClubQuery {
	return (&ClubapplicationClient{config: c.config}).QueryClub(c)
}

// Update returns a builder for updating this Clubapplication.
// Note that, you need to call Clubapplication.Unwrap() before calling this method, if this Clubapplication
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Clubapplication) Update() *ClubapplicationUpdateOne {
	return (&ClubapplicationClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Clubapplication) Unwrap() *Clubapplication {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Clubapplication is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Clubapplication) String() string {
	var builder strings.Builder
	builder.WriteString("Clubapplication(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", contact=")
	builder.WriteString(c.Contact)
	builder.WriteString(", reason=")
	builder.WriteString(c.Reason)
	builder.WriteString(", addeddatetime=")
	builder.WriteString(c.Addeddatetime.Format(time.ANSIC))
	builder.WriteString(", addername=")
	builder.WriteString(c.Addername)
	builder.WriteString(", discipline=")
	builder.WriteString(c.Discipline)
	builder.WriteString(", gender=")
	builder.WriteString(c.Gender)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", c.Age))
	builder.WriteString(", yaer=")
	builder.WriteString(fmt.Sprintf("%v", c.Yaer))
	builder.WriteByte(')')
	return builder.String()
}

// Clubapplications is a parsable slice of Clubapplication.
type Clubapplications []*Clubapplication

func (c Clubapplications) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
