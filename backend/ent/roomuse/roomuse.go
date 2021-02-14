// Code generated by entc, DO NOT EDIT.

package roomuse

const (
	// Label holds the string label denoting the roomuse type in the database.
	Label = "roomuse"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPeople holds the string denoting the people field in the database.
	FieldPeople = "people"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldContact holds the string denoting the contact field in the database.
	FieldContact = "contact"
	// FieldInTime holds the string denoting the in_time field in the database.
	FieldInTime = "in_time"
	// FieldOutTime holds the string denoting the out_time field in the database.
	FieldOutTime = "out_time"

	// EdgeRooms holds the string denoting the rooms edge name in mutations.
	EdgeRooms = "rooms"
	// EdgePurposes holds the string denoting the purposes edge name in mutations.
	EdgePurposes = "purposes"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"

	// Table holds the table name of the roomuse in the database.
	Table = "roomuses"
	// RoomsTable is the table the holds the rooms relation/edge.
	RoomsTable = "roomuses"
	// RoomsInverseTable is the table name for the Room entity.
	// It exists in this package in order to avoid circular dependency with the "room" package.
	RoomsInverseTable = "rooms"
	// RoomsColumn is the table column denoting the rooms relation/edge.
	RoomsColumn = "room_id"
	// PurposesTable is the table the holds the purposes relation/edge.
	PurposesTable = "roomuses"
	// PurposesInverseTable is the table name for the Purpose entity.
	// It exists in this package in order to avoid circular dependency with the "purpose" package.
	PurposesInverseTable = "purposes"
	// PurposesColumn is the table column denoting the purposes relation/edge.
	PurposesColumn = "purpose_id"
	// UsersTable is the table the holds the users relation/edge.
	UsersTable = "roomuses"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "UserID"
)

// Columns holds all SQL columns for roomuse fields.
var Columns = []string{
	FieldID,
	FieldPeople,
	FieldNote,
	FieldContact,
	FieldInTime,
	FieldOutTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Roomuse type.
var ForeignKeys = []string{
	"purpose_id",
	"room_id",
	"UserID",
}

var (
	// PeopleValidator is a validator for the "people" field. It is called by the builders before save.
	PeopleValidator func(int) error
	// NoteValidator is a validator for the "note" field. It is called by the builders before save.
	NoteValidator func(string) error
	// ContactValidator is a validator for the "contact" field. It is called by the builders before save.
	ContactValidator func(string) error
)
