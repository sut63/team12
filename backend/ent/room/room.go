// Code generated by entc, DO NOT EDIT.

package room

const (
	// Label holds the string label denoting the room type in the database.
	Label = "room"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRoomName holds the string denoting the room_name field in the database.
	FieldRoomName = "room_name"
	// FieldRoomLocation holds the string denoting the room_location field in the database.
	FieldRoomLocation = "room_location"
	// FieldRoomFloor holds the string denoting the room_floor field in the database.
	FieldRoomFloor = "room_floor"
	// FieldMaxContain holds the string denoting the max_contain field in the database.
	FieldMaxContain = "max_contain"

	// EdgeRoomuses holds the string denoting the roomuses edge name in mutations.
	EdgeRoomuses = "roomuses"

	// Table holds the table name of the room in the database.
	Table = "rooms"
	// RoomusesTable is the table the holds the roomuses relation/edge.
	RoomusesTable = "roomuses"
	// RoomusesInverseTable is the table name for the Roomuse entity.
	// It exists in this package in order to avoid circular dependency with the "roomuse" package.
	RoomusesInverseTable = "roomuses"
	// RoomusesColumn is the table column denoting the roomuses relation/edge.
	RoomusesColumn = "room_id"
)

// Columns holds all SQL columns for room fields.
var Columns = []string{
	FieldID,
	FieldRoomName,
	FieldRoomLocation,
	FieldRoomFloor,
	FieldMaxContain,
}

var (
	// RoomNameValidator is a validator for the "room_name" field. It is called by the builders before save.
	RoomNameValidator func(string) error
	// RoomLocationValidator is a validator for the "room_location" field. It is called by the builders before save.
	RoomLocationValidator func(string) error
	// RoomFloorValidator is a validator for the "room_floor" field. It is called by the builders before save.
	RoomFloorValidator func(string) error
	// MaxContainValidator is a validator for the "max_contain" field. It is called by the builders before save.
	MaxContainValidator func(int) error
)
