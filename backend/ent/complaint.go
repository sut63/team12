// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/OMENX/app/ent/club"
	"github.com/OMENX/app/ent/complaint"
	"github.com/OMENX/app/ent/complainttype"
	"github.com/OMENX/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Complaint is the model entity for the Complaint schema.
type Complaint struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Info holds the value of the "info" field.
	Info string `json:"info,omitempty"`
	// Date holds the value of the "date" field.
	Date time.Time `json:"date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ComplaintQuery when eager-loading is set.
	Edges  ComplaintEdges `json:"edges"`
	ClubID *int
	Type   *int
	UserID *int
}

// ComplaintEdges holds the relations/edges for other nodes in the graph.
type ComplaintEdges struct {
	// ComplaintToUser holds the value of the ComplaintToUser edge.
	ComplaintToUser *User
	// ComplaintToComplaintType holds the value of the ComplaintToComplaintType edge.
	ComplaintToComplaintType *ComplaintType
	// ComplaintToClub holds the value of the ComplaintToClub edge.
	ComplaintToClub *Club
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ComplaintToUserOrErr returns the ComplaintToUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ComplaintEdges) ComplaintToUserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.ComplaintToUser == nil {
			// The edge ComplaintToUser was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.ComplaintToUser, nil
	}
	return nil, &NotLoadedError{edge: "ComplaintToUser"}
}

// ComplaintToComplaintTypeOrErr returns the ComplaintToComplaintType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ComplaintEdges) ComplaintToComplaintTypeOrErr() (*ComplaintType, error) {
	if e.loadedTypes[1] {
		if e.ComplaintToComplaintType == nil {
			// The edge ComplaintToComplaintType was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: complainttype.Label}
		}
		return e.ComplaintToComplaintType, nil
	}
	return nil, &NotLoadedError{edge: "ComplaintToComplaintType"}
}

// ComplaintToClubOrErr returns the ComplaintToClub value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ComplaintEdges) ComplaintToClubOrErr() (*Club, error) {
	if e.loadedTypes[2] {
		if e.ComplaintToClub == nil {
			// The edge ComplaintToClub was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: club.Label}
		}
		return e.ComplaintToClub, nil
	}
	return nil, &NotLoadedError{edge: "ComplaintToClub"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Complaint) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // info
		&sql.NullTime{},   // date
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Complaint) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // ClubID
		&sql.NullInt64{}, // Type
		&sql.NullInt64{}, // UserID
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Complaint fields.
func (c *Complaint) assignValues(values ...interface{}) error {
	if m, n := len(values), len(complaint.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field info", values[0])
	} else if value.Valid {
		c.Info = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field date", values[1])
	} else if value.Valid {
		c.Date = value.Time
	}
	values = values[2:]
	if len(values) == len(complaint.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field ClubID", value)
		} else if value.Valid {
			c.ClubID = new(int)
			*c.ClubID = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field Type", value)
		} else if value.Valid {
			c.Type = new(int)
			*c.Type = int(value.Int64)
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

// QueryComplaintToUser queries the ComplaintToUser edge of the Complaint.
func (c *Complaint) QueryComplaintToUser() *UserQuery {
	return (&ComplaintClient{config: c.config}).QueryComplaintToUser(c)
}

// QueryComplaintToComplaintType queries the ComplaintToComplaintType edge of the Complaint.
func (c *Complaint) QueryComplaintToComplaintType() *ComplaintTypeQuery {
	return (&ComplaintClient{config: c.config}).QueryComplaintToComplaintType(c)
}

// QueryComplaintToClub queries the ComplaintToClub edge of the Complaint.
func (c *Complaint) QueryComplaintToClub() *ClubQuery {
	return (&ComplaintClient{config: c.config}).QueryComplaintToClub(c)
}

// Update returns a builder for updating this Complaint.
// Note that, you need to call Complaint.Unwrap() before calling this method, if this Complaint
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Complaint) Update() *ComplaintUpdateOne {
	return (&ComplaintClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Complaint) Unwrap() *Complaint {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Complaint is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Complaint) String() string {
	var builder strings.Builder
	builder.WriteString("Complaint(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", info=")
	builder.WriteString(c.Info)
	builder.WriteString(", date=")
	builder.WriteString(c.Date.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Complaints is a parsable slice of Complaint.
type Complaints []*Complaint

func (c Complaints) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}