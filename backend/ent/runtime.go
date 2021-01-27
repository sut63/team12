// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/OMENX/app/ent/academicyear"
	"github.com/OMENX/app/ent/activities"
	"github.com/OMENX/app/ent/activitytype"
	"github.com/OMENX/app/ent/club"
	"github.com/OMENX/app/ent/clubapplication"
	"github.com/OMENX/app/ent/clubbranch"
	"github.com/OMENX/app/ent/clubtype"
	"github.com/OMENX/app/ent/complaint"
	"github.com/OMENX/app/ent/purpose"
	"github.com/OMENX/app/ent/room"
	"github.com/OMENX/app/ent/schema"
	"github.com/OMENX/app/ent/user"
	"github.com/OMENX/app/ent/usertype"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	academicyearFields := schema.AcademicYear{}.Fields()
	_ = academicyearFields
	// academicyearDescSemester is the schema descriptor for semester field.
	academicyearDescSemester := academicyearFields[0].Descriptor()
	// academicyear.SemesterValidator is a validator for the "semester" field. It is called by the builders before save.
	academicyear.SemesterValidator = academicyearDescSemester.Validators[0].(func(string) error)
	activitiesFields := schema.Activities{}.Fields()
	_ = activitiesFields
	// activitiesDescName is the schema descriptor for name field.
	activitiesDescName := activitiesFields[0].Descriptor()
	// activities.NameValidator is a validator for the "name" field. It is called by the builders before save.
	activities.NameValidator = activitiesDescName.Validators[0].(func(string) error)
	// activitiesDescDetail is the schema descriptor for detail field.
	activitiesDescDetail := activitiesFields[1].Descriptor()
	// activities.DetailValidator is a validator for the "detail" field. It is called by the builders before save.
	activities.DetailValidator = activitiesDescDetail.Validators[0].(func(string) error)
	// activitiesDescLocation is the schema descriptor for location field.
	activitiesDescLocation := activitiesFields[2].Descriptor()
	// activities.LocationValidator is a validator for the "location" field. It is called by the builders before save.
	activities.LocationValidator = activitiesDescLocation.Validators[0].(func(string) error)
	activitytypeFields := schema.ActivityType{}.Fields()
	_ = activitytypeFields
	// activitytypeDescName is the schema descriptor for name field.
	activitytypeDescName := activitytypeFields[0].Descriptor()
	// activitytype.NameValidator is a validator for the "name" field. It is called by the builders before save.
	activitytype.NameValidator = activitytypeDescName.Validators[0].(func(string) error)
	clubFields := schema.Club{}.Fields()
	_ = clubFields
	// clubDescName is the schema descriptor for name field.
	clubDescName := clubFields[0].Descriptor()
	// club.NameValidator is a validator for the "name" field. It is called by the builders before save.
	club.NameValidator = clubDescName.Validators[0].(func(string) error)
	// clubDescPurpose is the schema descriptor for purpose field.
	clubDescPurpose := clubFields[1].Descriptor()
	// club.PurposeValidator is a validator for the "purpose" field. It is called by the builders before save.
	club.PurposeValidator = clubDescPurpose.Validators[0].(func(string) error)
	clubbranchFields := schema.ClubBranch{}.Fields()
	_ = clubbranchFields
	// clubbranchDescName is the schema descriptor for name field.
	clubbranchDescName := clubbranchFields[0].Descriptor()
	// clubbranch.NameValidator is a validator for the "name" field. It is called by the builders before save.
	clubbranch.NameValidator = clubbranchDescName.Validators[0].(func(string) error)
	clubtypeFields := schema.ClubType{}.Fields()
	_ = clubtypeFields
	// clubtypeDescName is the schema descriptor for name field.
	clubtypeDescName := clubtypeFields[0].Descriptor()
	// clubtype.NameValidator is a validator for the "name" field. It is called by the builders before save.
	clubtype.NameValidator = clubtypeDescName.Validators[0].(func(string) error)
	clubapplicationFields := schema.Clubapplication{}.Fields()
	_ = clubapplicationFields
	// clubapplicationDescContact is the schema descriptor for contact field.
	clubapplicationDescContact := clubapplicationFields[0].Descriptor()
	// clubapplication.ContactValidator is a validator for the "contact" field. It is called by the builders before save.
	clubapplication.ContactValidator = clubapplicationDescContact.Validators[0].(func(string) error)
	// clubapplicationDescAddername is the schema descriptor for addername field.
	clubapplicationDescAddername := clubapplicationFields[3].Descriptor()
	// clubapplication.AddernameValidator is a validator for the "addername" field. It is called by the builders before save.
	clubapplication.AddernameValidator = clubapplicationDescAddername.Validators[0].(func(string) error)
	// clubapplicationDescDiscipline is the schema descriptor for discipline field.
	clubapplicationDescDiscipline := clubapplicationFields[4].Descriptor()
	// clubapplication.DisciplineValidator is a validator for the "discipline" field. It is called by the builders before save.
	clubapplication.DisciplineValidator = clubapplicationDescDiscipline.Validators[0].(func(string) error)
	// clubapplicationDescGender is the schema descriptor for gender field.
	clubapplicationDescGender := clubapplicationFields[5].Descriptor()
	// clubapplication.GenderValidator is a validator for the "gender" field. It is called by the builders before save.
	clubapplication.GenderValidator = clubapplicationDescGender.Validators[0].(func(string) error)
	// clubapplicationDescAge is the schema descriptor for age field.
	clubapplicationDescAge := clubapplicationFields[6].Descriptor()
	// clubapplication.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	clubapplication.AgeValidator = func() func(int) error {
		validators := clubapplicationDescAge.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(age int) error {
			for _, fn := range fns {
				if err := fn(age); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// clubapplicationDescYaer is the schema descriptor for yaer field.
	clubapplicationDescYaer := clubapplicationFields[7].Descriptor()
	// clubapplication.YaerValidator is a validator for the "yaer" field. It is called by the builders before save.
	clubapplication.YaerValidator = clubapplicationDescYaer.Validators[0].(func(int) error)
	complaintFields := schema.Complaint{}.Fields()
	_ = complaintFields
	// complaintDescName is the schema descriptor for name field.
	complaintDescName := complaintFields[0].Descriptor()
	// complaint.NameValidator is a validator for the "name" field. It is called by the builders before save.
	complaint.NameValidator = complaintDescName.Validators[0].(func(string) error)
	// complaintDescPhonenumber is the schema descriptor for phonenumber field.
	complaintDescPhonenumber := complaintFields[1].Descriptor()
	// complaint.PhonenumberValidator is a validator for the "phonenumber" field. It is called by the builders before save.
	complaint.PhonenumberValidator = func() func(string) error {
		validators := complaintDescPhonenumber.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(phonenumber string) error {
			for _, fn := range fns {
				if err := fn(phonenumber); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// complaintDescInfo is the schema descriptor for info field.
	complaintDescInfo := complaintFields[2].Descriptor()
	// complaint.InfoValidator is a validator for the "info" field. It is called by the builders before save.
	complaint.InfoValidator = complaintDescInfo.Validators[0].(func(string) error)
	purposeFields := schema.Purpose{}.Fields()
	_ = purposeFields
	// purposeDescPurpose is the schema descriptor for purpose field.
	purposeDescPurpose := purposeFields[0].Descriptor()
	// purpose.PurposeValidator is a validator for the "purpose" field. It is called by the builders before save.
	purpose.PurposeValidator = purposeDescPurpose.Validators[0].(func(string) error)
	roomFields := schema.Room{}.Fields()
	_ = roomFields
	// roomDescRoomName is the schema descriptor for room_name field.
	roomDescRoomName := roomFields[0].Descriptor()
	// room.RoomNameValidator is a validator for the "room_name" field. It is called by the builders before save.
	room.RoomNameValidator = roomDescRoomName.Validators[0].(func(string) error)
	// roomDescRoomLocation is the schema descriptor for room_location field.
	roomDescRoomLocation := roomFields[1].Descriptor()
	// room.RoomLocationValidator is a validator for the "room_location" field. It is called by the builders before save.
	room.RoomLocationValidator = roomDescRoomLocation.Validators[0].(func(string) error)
	// roomDescRoomFloor is the schema descriptor for room_floor field.
	roomDescRoomFloor := roomFields[2].Descriptor()
	// room.RoomFloorValidator is a validator for the "room_floor" field. It is called by the builders before save.
	room.RoomFloorValidator = roomDescRoomFloor.Validators[0].(func(string) error)
	// roomDescMaxContain is the schema descriptor for max_contain field.
	roomDescMaxContain := roomFields[3].Descriptor()
	// room.MaxContainValidator is a validator for the "max_contain" field. It is called by the builders before save.
	room.MaxContainValidator = roomDescMaxContain.Validators[0].(func(int) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[1].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	usertypeFields := schema.Usertype{}.Fields()
	_ = usertypeFields
	// usertypeDescName is the schema descriptor for name field.
	usertypeDescName := usertypeFields[0].Descriptor()
	// usertype.NameValidator is a validator for the "name" field. It is called by the builders before save.
	usertype.NameValidator = usertypeDescName.Validators[0].(func(string) error)
}
