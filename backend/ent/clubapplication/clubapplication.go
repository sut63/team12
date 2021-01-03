// Code generated by entc, DO NOT EDIT.

package clubapplication

import (
	"time"
)

const (
	// Label holds the string label denoting the clubapplication type in the database.
	Label = "clubapplication"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldApplyname holds the string denoting the applyname field in the database.
	FieldApplyname = "applyname"
	// FieldContact holds the string denoting the contact field in the database.
	FieldContact = "contact"
	// FieldReason holds the string denoting the reason field in the database.
	FieldReason = "reason"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeClubappstatus holds the string denoting the clubappstatus edge name in mutations.
	EdgeClubappstatus = "clubappstatus"
	// EdgeClub holds the string denoting the club edge name in mutations.
	EdgeClub = "club"

	// Table holds the table name of the clubapplication in the database.
	Table = "clubapplications"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "clubapplications"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "UserID"
	// ClubappstatusTable is the table the holds the clubappstatus relation/edge.
	ClubappstatusTable = "clubapplications"
	// ClubappstatusInverseTable is the table name for the ClubappStatus entity.
	// It exists in this package in order to avoid circular dependency with the "clubappstatus" package.
	ClubappstatusInverseTable = "clubapp_status"
	// ClubappstatusColumn is the table column denoting the clubappstatus relation/edge.
	ClubappstatusColumn = "clubappstatus_id"
	// ClubTable is the table the holds the club relation/edge.
	ClubTable = "clubapplications"
	// ClubInverseTable is the table name for the Club entity.
	// It exists in this package in order to avoid circular dependency with the "club" package.
	ClubInverseTable = "clubs"
	// ClubColumn is the table column denoting the club relation/edge.
	ClubColumn = "ClubID"
)

// Columns holds all SQL columns for clubapplication fields.
var Columns = []string{
	FieldID,
	FieldApplyname,
	FieldContact,
	FieldReason,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Clubapplication type.
var ForeignKeys = []string{
	"ClubID",
	"clubappstatus_id",
	"UserID",
}

var (
	// ApplynameValidator is a validator for the "applyname" field. It is called by the builders before save.
	ApplynameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
)
