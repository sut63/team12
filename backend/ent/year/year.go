// Code generated by entc, DO NOT EDIT.

package year

const (
	// Label holds the string label denoting the year type in the database.
	Label = "year"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"

	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"

	// Table holds the table name of the year in the database.
	Table = "years"
	// UsersTable is the table the holds the users relation/edge.
	UsersTable = "users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "year_id"
)

// Columns holds all SQL columns for year fields.
var Columns = []string{
	FieldID,
	FieldYear,
}