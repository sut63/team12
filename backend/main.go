package main

import (
	"context"
	"fmt"
	"log"

	"github.com/OMENX/app/controllers"
	_ "github.com/OMENX/app/docs"
	"github.com/OMENX/app/ent"
	"github.com/OMENX/app/ent/club"
	"github.com/OMENX/app/ent/usertype"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Users struct {
	User []User
}

type User struct {
	name     string
	email    string
	password string
	usertype int
	club     int
}

type Usertypes struct {
	Usertype []Usertype
}

type Usertype struct {
	name string
}

type ClubBranchs struct {
	ClubBranch []ClubBranch
}

type ClubBranch struct {
	name string
}
type ClubTypes struct {
	ClubType []ClubType
}

type ClubType struct {
	name string
}

type ActivityTypes struct {
	ActivityType []ActivityType
}

type ActivityType struct {
	name string
}

type AcademicYears struct {
	AcademicYear []AcademicYear
}

type AcademicYear struct {
	semester string
}

// Disciplines is ...
type Disciplines struct {
	Discipline []Discipline
}

// Discipline is ...
type Discipline struct {
	Discip string
}

// Years is ...
type Years struct {
	Year []Year
}

// Year is ...
type Year struct {
	Year int
}

// Genders is ...
type Genders struct {
	Gender []Gender
}

// Gender is ...
type Gender struct {
	Gender string
}

// UserStatuss is ...
type UserStatuss struct {
	UserStatus []UserStatus
}

// UserStatus is ...
type UserStatus struct {
	UserStatus string
}

// ClubappStatuses is ...
type ClubappStatuses struct {
	ClubappStatus []ClubappStatus
}

// ClubappStatus is ...
type ClubappStatus struct {
	Status string
}

// Complaint struct
type Complaint struct {
	UserID int
	ClubID int
	TypeID int
	Info   string
	Date   string
}

// Complaints struct
type Complaints struct {
	Complaint []Complaint
}

// ComplaintType struct
type ComplaintType struct {
	Description string
}

// ComplaintTypes struct
type ComplaintTypes struct {
	ComplaintType []ComplaintType
}

type Rooms struct {
	Room []Room
}

type Room struct {
	ROOMNAME     string
	ROOMLOCATION string
	ROOMFLOOR    string
	MAXCONTAIN   int
}

type Purposes struct {
	Purpose []Purpose
}

type Purpose struct {
	PURPOSE string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewAcademicYearController(v1, client)
	controllers.NewActivitiesController(v1, client)
	controllers.NewActivityTypeController(v1, client)
	controllers.NewUsertypeController(v1, client)

	controllers.NewDisciplineController(v1, client)
	controllers.NewYearController(v1, client)
	controllers.NewGenderController(v1, client)
	controllers.NewUserstatusController(v1, client)
	controllers.NewClubappStatusController(v1, client)
	controllers.NewClubapplicationController(v1, client)
	controllers.NewComplaintController(v1, client)
	controllers.NewComplainttypeController(v1, client)
	controllers.NewClubBranchController(v1, client)
	controllers.NewClubTypeController(v1, client)
	controllers.NewClubController(v1, client)
	//เพิ่มล่าสุด
	controllers.NewRoomController(v1, client)
	controllers.NewPurposeController(v1, client)
	controllers.NewRoomuseController(v1, client)

	// Set Types Data
	typedata := Usertypes{
		Usertype: []Usertype{
			Usertype{"นักศึกษา"},
			Usertype{"เจ้าหน้าที่"},
		},
	}

	for _, t := range typedata.Usertype {
		client.Usertype.
			Create().
			SetName(t.name).
			Save(context.Background())
	}

		// Set ClubBranch Data
	ClubBranchs := ClubBranchs{
		ClubBranch: []ClubBranch{
			ClubBranch{"B4101"},
			ClubBranch{"B1212"},
			ClubBranch{"สนามแบตมินตัน"},
		},
	}

	for _, cb := range ClubBranchs.ClubBranch {
		client.ClubBranch.
			Create().
			SetName(cb.name).
			Save(context.Background())
	}
		// Set ClubType Data
	ClubTypes := ClubTypes{
		ClubType: []ClubType{
			ClubType{"กีฬา"},
			ClubType{"วิชาการ"},
			ClubType{"แข่งขัน"},
		},
	}

	for _, ct := range ClubTypes.ClubType {
		client.ClubType.
			Create().
			SetName(ct.name).
			Save(context.Background())
	}

	// Set Users Data
	users := Users{
		User: []User{
			User{"Suphasin", "Suphasin@gmail.com", "123456789A", 1, 2},
			User{"Yosang", "Yosang@gmail.com", "0000FUSE", 2, 1},
		},
	}

	for _, u := range users.User {

		t, err := client.Usertype.
			Query().
			Where(usertype.IDEQ(int(u.usertype))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		c, err := client.Club.
			Query().
			Where(club.IDEQ(int(u.club))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.User.
			Create().
			SetName(u.name).
			SetEmail(u.email).
			SetPassword(u.password).
			SetUsertype(t).
			SetClubuser(c).
			Save(context.Background())
	}

	// Set ActivityType Data
	activitytypes := ActivityTypes{
		ActivityType: []ActivityType{
			ActivityType{"พัฒนาสังคม"},
			ActivityType{"การศึกษา"},
			ActivityType{"สิ่งแวดล้อม"},
		},
	}

	for _, act := range activitytypes.ActivityType {
		client.ActivityType.
			Create().
			SetName(act.name).
			Save(context.Background())
	}

	// Set AcademicYear Data
	academicyears := AcademicYears{
		AcademicYear: []AcademicYear{
			AcademicYear{"1/2563"},
			AcademicYear{"2/2563"},
			AcademicYear{"3/2563"},
		},
	}

	for _, acy := range academicyears.AcademicYear {
		client.AcademicYear.
			Create().
			SetSemester(acy.semester).
			Save(context.Background())
	}

	// Set ComplaintType Data
	complainttypes := ComplaintTypes{
		ComplaintType: []ComplaintType{
			ComplaintType{"TestType1"},
			ComplaintType{"TestType2"},
		},
	}

	for _, cpt := range complainttypes.ComplaintType {
		client.ComplaintType.
			Create().
			SetDescription(cpt.Description).
			Save(context.Background())
	}

	// Set Complaint Data
	complaints := Complaints{
		Complaint: []Complaint{
			Complaint{1, 1, 1, "Test1", "Test1"},
			Complaint{2, 2, 2, "Test2", "Test2"},
		},
	}

	for _, cp := range complaints.Complaint {
		client.Complaint.
			Create().
			SetComplaintToUserID(cp.UserID).
			SetComplaintToClubID(cp.ClubID).
			SetComplaintToComplaintTypeID(cp.TypeID).
			SetInfo(cp.Info).
			SetDate(cp.Date).
			Save(context.Background())
	}

	// Set Room Data
	rooms := Rooms{
		Room: []Room{
			Room{"Meeting room", "Building A", "floor G", 10},
			Room{"Relax room", "Building A", "floor 1", 20},
		},
	}

	for _, r := range rooms.Room {
		client.Room.
			Create().
			SetRoomName(r.ROOMNAME).
			SetRoomLocation(r.ROOMLOCATION).
			SetRoomFloor(r.ROOMFLOOR).
			SetMaxContain(r.MAXCONTAIN).
			Save(context.Background())
	}

	// Set Purpose Data
	purposes := Purposes{
		Purpose: []Purpose{
			Purpose{"ประชุม"},
			Purpose{"ทำชมรม"},
			Purpose{"จัดกิจกรรมชมรม"},
		},
	}

	for _, p := range purposes.Purpose {
		client.Purpose.
			Create().
			SetPurpose(p.PURPOSE).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
