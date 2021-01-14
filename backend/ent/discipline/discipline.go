// Code generated by entc, DO NOT EDIT.

package discipline

const (
	// Label holds the string label denoting the discipline type in the database.
	Label = "discipline"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDiscipline holds the string denoting the discipline field in the database.
	FieldDiscipline = "discipline"

	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"

	// Table holds the table name of the discipline in the database.
	Table = "disciplines"
	// UsersTable is the table the holds the users relation/edge.
	UsersTable = "users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "discipline_id"
)

// Columns holds all SQL columns for discipline fields.
var Columns = []string{
	FieldID,
	FieldDiscipline,
}