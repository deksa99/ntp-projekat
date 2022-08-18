package database

import (
	"AppointmentService/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var (
	requests = []model.AppointmentRequest{
		{UserID: 1, VehicleID: 1, ServiceID: 1, Status: model.Submitted, SubmittedAt: time.Date(2022, 8, 20, 19, 42, 6, 0, time.Local)},
		{UserID: 1, VehicleID: 2, ServiceID: 2, Status: model.Submitted, SubmittedAt: time.Date(2022, 8, 20, 19, 47, 34, 0, time.Local)},
		{UserID: 1, VehicleID: 2, ServiceID: 11, Status: model.Accepted, SubmittedAt: time.Date(2022, 7, 13, 17, 37, 3, 27, time.Local)},
		{UserID: 1, VehicleID: 2, ServiceID: 12, Status: model.Accepted, SubmittedAt: time.Date(2022, 7, 16, 17, 24, 59, 0, time.Local)},

		{UserID: 2, VehicleID: 3, ServiceID: 13, Status: model.Submitted, SubmittedAt: time.Date(2022, 8, 18, 10, 53, 13, 0, time.Local)},
		{UserID: 2, VehicleID: 3, ServiceID: 15, Status: model.Submitted, SubmittedAt: time.Date(2022, 8, 19, 12, 31, 11, 0, time.Local)},
		{UserID: 2, VehicleID: 3, ServiceID: 4, Status: model.Accepted, SubmittedAt: time.Date(2022, 7, 27, 12, 40, 33, 0, time.Local)},
		{UserID: 2, VehicleID: 3, ServiceID: 5, Status: model.Accepted, SubmittedAt: time.Date(2022, 8, 17, 15, 12, 37, 0, time.Local)},
		{UserID: 2, VehicleID: 3, ServiceID: 23, Status: model.Accepted, SubmittedAt: time.Date(2022, 8, 14, 23, 20, 39, 0, time.Local)},
		{UserID: 2, VehicleID: 3, ServiceID: 24, Status: model.Rejected, SubmittedAt: time.Date(2022, 5, 11, 3, 37, 3, 0, time.Local)},

		{UserID: 3, VehicleID: 4, ServiceID: 4, Status: model.Accepted, SubmittedAt: time.Date(2022, 8, 18, 10, 23, 7, 0, time.Local)},
		{UserID: 3, VehicleID: 4, ServiceID: 19, Status: model.Rejected, SubmittedAt: time.Date(2022, 8, 14, 9, 40, 32, 0, time.Local)},
		{UserID: 3, VehicleID: 4, ServiceID: 19, Status: model.Rejected, SubmittedAt: time.Date(2022, 8, 15, 9, 40, 31, 0, time.Local)},
	}
	appointments = []model.Appointment{
		{UserID: 1, VehicleID: 2, ServiceID: 11, WorkerID: 3, StartTime: time.Date(2022, 7, 15, 10, 0, 0, 0, time.Local), FinishTime: time.Date(2022, 7, 16, 14, 34, 44, 0, time.Local), Status: model.Finished},
		{UserID: 1, VehicleID: 2, ServiceID: 12, WorkerID: 4, StartTime: time.Date(2022, 7, 20, 10, 0, 0, 0, time.Local), FinishTime: time.Date(2022, 7, 23, 15, 27, 49, 0, time.Local), Status: model.Finished},

		{UserID: 2, VehicleID: 3, ServiceID: 4, WorkerID: 1, StartTime: time.Date(2022, 7, 29, 12, 30, 0, 0, time.Local), FinishTime: time.Date(2022, 8, 1, 12, 30, 27, 0, time.Local), Status: model.Finished},
		{UserID: 2, VehicleID: 3, ServiceID: 5, WorkerID: 2, StartTime: time.Date(2022, 8, 25, 15, 0, 0, 0, time.Local), FinishTime: time.Time{}, Status: model.Scheduled},
		{UserID: 2, VehicleID: 3, ServiceID: 23, WorkerID: 6, StartTime: time.Date(2022, 8, 16, 14, 0, 0, 0, time.Local), FinishTime: time.Time{}, Status: model.Cancelled},

		{UserID: 3, VehicleID: 4, ServiceID: 4, WorkerID: 1, StartTime: time.Date(2022, 8, 26, 9, 0, 0, 0, time.Local), FinishTime: time.Time{}, Status: model.Scheduled},
	}
)

var Db *gorm.DB
var err error

func ConnectToDatabase() {
	connStr := os.Getenv("DB_CONN_STR")
	Db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully connected to DB")
	}

	err = Db.Migrator().DropTable(&model.AppointmentRequest{})
	if err != nil {
		log.Panic(err.Error())
	}
	err = Db.Migrator().DropTable(&model.Appointment{})
	if err != nil {
		log.Panic(err.Error())
	}

	log.Println("Tables dropped successfully")

	err = Db.AutoMigrate(&model.AppointmentRequest{})
	if err != nil {
		log.Panic(err.Error())
	}
	err = Db.AutoMigrate(&model.Appointment{})
	if err != nil {
		log.Panic(err.Error())
	}

	log.Println("Migration successful")

	for _, r := range requests {
		Db.Create(&r)
	}
	for _, a := range appointments {
		Db.Create(&a)
	}

	log.Println("DB seeding successful")
}
