// Code generated by entc, DO NOT EDIT.

package complaint

const (
	// Label holds the string label denoting the complaint type in the database.
	Label = "complaint"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldInfo holds the string denoting the info field in the database.
	FieldInfo = "info"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"

	// EdgeComplaintToUser holds the string denoting the complainttouser edge name in mutations.
	EdgeComplaintToUser = "ComplaintToUser"
	// EdgeComplaintToComplaintType holds the string denoting the complainttocomplainttype edge name in mutations.
	EdgeComplaintToComplaintType = "ComplaintToComplaintType"
	// EdgeComplaintToClub holds the string denoting the complainttoclub edge name in mutations.
	EdgeComplaintToClub = "ComplaintToClub"

	// Table holds the table name of the complaint in the database.
	Table = "complaints"
	// ComplaintToUserTable is the table the holds the ComplaintToUser relation/edge.
	ComplaintToUserTable = "complaints"
	// ComplaintToUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ComplaintToUserInverseTable = "users"
	// ComplaintToUserColumn is the table column denoting the ComplaintToUser relation/edge.
	ComplaintToUserColumn = "UserID"
	// ComplaintToComplaintTypeTable is the table the holds the ComplaintToComplaintType relation/edge.
	ComplaintToComplaintTypeTable = "complaints"
	// ComplaintToComplaintTypeInverseTable is the table name for the ComplaintType entity.
	// It exists in this package in order to avoid circular dependency with the "complainttype" package.
	ComplaintToComplaintTypeInverseTable = "complaint_types"
	// ComplaintToComplaintTypeColumn is the table column denoting the ComplaintToComplaintType relation/edge.
	ComplaintToComplaintTypeColumn = "Type"
	// ComplaintToClubTable is the table the holds the ComplaintToClub relation/edge.
	ComplaintToClubTable = "complaints"
	// ComplaintToClubInverseTable is the table name for the Club entity.
	// It exists in this package in order to avoid circular dependency with the "club" package.
	ComplaintToClubInverseTable = "clubs"
	// ComplaintToClubColumn is the table column denoting the ComplaintToClub relation/edge.
	ComplaintToClubColumn = "ClubID"
)

// Columns holds all SQL columns for complaint fields.
var Columns = []string{
	FieldID,
	FieldInfo,
	FieldDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Complaint type.
var ForeignKeys = []string{
	"ClubID",
	"Type",
	"UserID",
}
