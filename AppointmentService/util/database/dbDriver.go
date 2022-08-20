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
		{UserID: 1, FirstName: "Mika", LastName: "Mikic", Email: "user1@maildrop.cc",
			VehicleID: 1, Manufacturer: "Honda", CarModel: "Civic", LicencePlate: "NS123ABC", ChassisNumber: "ABC123456ABC",
			ServiceID: 1, ServiceName: "Mali servis", ServicePrice: 7300, Available: true,
			CarServiceID: 1, CarServiceName: "Kvačilo", Street: "Bulevar oslobođenja 21, Novi Sad",
			Status: model.Submitted, SubmittedAt: time.Date(2022, 8, 20, 19, 42, 6, 0, time.Local)},

		{UserID: 1, FirstName: "Mika", LastName: "Mikic", Email: "user1@maildrop.cc",
			VehicleID: 2, Manufacturer: "Chevrolet", CarModel: "Camaro", LicencePlate: "BG123BG", ChassisNumber: "CBA654321ABC",
			ServiceID: 2, ServiceName: "Veliki servis", ServicePrice: 27600, Available: true,
			CarServiceID: 1, CarServiceName: "Kvačilo", Street: "Bulevar oslobođenja 21, Novi Sad",
			Status: model.Submitted, SubmittedAt: time.Date(2022, 8, 20, 19, 47, 34, 0, time.Local)},

		{UserID: 1, FirstName: "Mika", LastName: "Mikic", Email: "user1@maildrop.cc",
			VehicleID: 2, Manufacturer: "Chevrolet", CarModel: "Camaro", LicencePlate: "BG123BG", ChassisNumber: "CBA654321ABC",
			ServiceID: 11, ServiceName: "Montaža gume", ServicePrice: 620, Available: true,
			CarServiceID: 2, CarServiceName: "Mile i sin", Street: "Vladimira Nazora 2, Petrovaradin",
			Status: model.Accepted, SubmittedAt: time.Date(2022, 7, 13, 17, 37, 3, 27, time.Local)},

		{UserID: 1, FirstName: "Mika", LastName: "Mikic", Email: "user1@maildrop.cc",
			VehicleID: 2, Manufacturer: "Chevrolet", CarModel: "Camaro", LicencePlate: "BG123BG", ChassisNumber: "CBA654321ABC",
			ServiceID: 12, ServiceName: "Balansiranje", ServicePrice: 880, Available: true,
			CarServiceID: 2, CarServiceName: "Mile i sin", Street: "Vladimira Nazora 2, Petrovaradin",
			Status: model.Accepted, SubmittedAt: time.Date(2022, 7, 16, 17, 24, 59, 0, time.Local)},
		//********************************************************

		{UserID: 2, FirstName: "Vera", LastName: "Veric", Email: "user2@maildrop.cc",
			VehicleID: 3, Manufacturer: "Ford", CarModel: "Mustang", LicencePlate: "NS555CBA", ChassisNumber: "AAA123456AAA",
			ServiceID: 13, ServiceName: "Mali servis", ServicePrice: 9700, Available: true,
			CarServiceID: 3, CarServiceName: "Auto Servis Roki", Street: "Veljka Petrovica 5, Novi Sad",
			Status: model.Submitted, SubmittedAt: time.Date(2022, 8, 18, 10, 53, 13, 0, time.Local)},

		{UserID: 2, FirstName: "Vera", LastName: "Veric", Email: "user2@maildrop.cc",
			VehicleID: 3, Manufacturer: "Ford", CarModel: "Mustang", LicencePlate: "NS555CBA", ChassisNumber: "AAA123456AAA",
			ServiceID: 15, ServiceName: "Skidanje točka", ServicePrice: 380, Available: true,
			CarServiceID: 3, CarServiceName: "Auto Servis Roki", Street: "Veljka Petrovica 5, Novi Sad",
			Status: model.Submitted, SubmittedAt: time.Date(2022, 8, 19, 12, 31, 11, 0, time.Local)},

		{UserID: 2, FirstName: "Vera", LastName: "Veric", Email: "user2@maildrop.cc",
			VehicleID: 3, Manufacturer: "Ford", CarModel: "Mustang", LicencePlate: "NS555CBA", ChassisNumber: "AAA123456AAA",
			ServiceID: 4, ServiceName: "Zamena sijalica", ServicePrice: 850, Available: true,
			CarServiceID: 1, CarServiceName: "Kvačilo", Street: "Bulevar oslobođenja 21, Novi Sad",
			Status: model.Accepted, SubmittedAt: time.Date(2022, 7, 27, 12, 40, 33, 0, time.Local)},

		{UserID: 2, FirstName: "Vera", LastName: "Veric", Email: "user2@maildrop.cc",
			VehicleID: 3, Manufacturer: "Ford", CarModel: "Mustang", LicencePlate: "NS555CBA", ChassisNumber: "AAA123456AAA",
			ServiceID: 5, ServiceName: "Zamena akulumatora", ServicePrice: 1350, Available: true,
			CarServiceID: 1, CarServiceName: "Kvačilo", Street: "Bulevar oslobođenja 21, Novi Sad",
			Status: model.Accepted, SubmittedAt: time.Date(2022, 8, 17, 15, 12, 37, 0, time.Local)},

		{UserID: 2, FirstName: "Vera", LastName: "Veric", Email: "user2@maildrop.cc",
			VehicleID: 3, Manufacturer: "Ford", CarModel: "Mustang", LicencePlate: "NS555CBA", ChassisNumber: "AAA123456AAA",
			ServiceID: 23, ServiceName: "Montaža gume", ServicePrice: 880, Available: true,
			CarServiceID: 4, CarServiceName: "Auto Servis Roki 2", Street: "Moše Pijade 17, Sremska Kamenica",
			Status: model.Accepted, SubmittedAt: time.Date(2022, 8, 14, 23, 20, 39, 0, time.Local)},

		{UserID: 2, FirstName: "Vera", LastName: "Veric", Email: "user2@maildrop.cc",
			VehicleID: 3, Manufacturer: "Ford", CarModel: "Mustang", LicencePlate: "NS555CBA", ChassisNumber: "AAA123456AAA",
			ServiceID: 24, ServiceName: "Balansiranje", ServicePrice: 1430, Available: true,
			CarServiceID: 4, CarServiceName: "Auto Servis Roki 2", Street: "Moše Pijade 17, Sremska Kamenica",
			Status: model.Rejected, SubmittedAt: time.Date(2022, 5, 11, 3, 37, 3, 0, time.Local)},
		//********************************************************

		{UserID: 3, FirstName: "Laza", LastName: "Lazic", Email: "user3@maildrop.cc",
			VehicleID: 4, Manufacturer: "Porsche", CarModel: "911 Turbo S", LicencePlate: "NS123456", ChassisNumber: "CCC123456CCC",
			ServiceID: 4, ServiceName: "Zamena sijalica", ServicePrice: 850, Available: true,
			CarServiceID: 1, CarServiceName: "Kvačilo", Street: "Bulevar oslobođenja 21, Novi Sad",
			Status: model.Accepted, SubmittedAt: time.Date(2022, 8, 18, 10, 23, 7, 0, time.Local)},

		{UserID: 3, FirstName: "Laza", LastName: "Lazic", Email: "user3@maildrop.cc",
			VehicleID: 4, Manufacturer: "Porsche", CarModel: "911 Turbo S", LicencePlate: "NS123456", ChassisNumber: "CCC123456CCC",
			ServiceID: 19, ServiceName: "Veliki servis", ServicePrice: 35050, Available: true,
			CarServiceID: 4, CarServiceName: "Auto Servis Roki 2", Street: "Moše Pijade 17, Sremska Kamenica",
			Status: model.CancelledRequest, SubmittedAt: time.Date(2022, 8, 1, 9, 40, 32, 0, time.Local)},

		{UserID: 3, FirstName: "Laza", LastName: "Lazic", Email: "user3@maildrop.cc",
			VehicleID: 4, Manufacturer: "Porsche", CarModel: "911 Turbo S", LicencePlate: "NS123456", ChassisNumber: "CCC123456CCC",
			ServiceID: 19, ServiceName: "Veliki servis", ServicePrice: 35050, Available: true,
			CarServiceID: 4, CarServiceName: "Auto Servis Roki 2", Street: "Moše Pijade 17, Sremska Kamenica",
			Status: model.Rejected, SubmittedAt: time.Date(2022, 8, 14, 9, 40, 32, 0, time.Local)},

		{UserID: 3, FirstName: "Laza", LastName: "Lazic", Email: "user3@maildrop.cc",
			VehicleID: 4, Manufacturer: "Porsche", CarModel: "911 Turbo S", LicencePlate: "NS123456", ChassisNumber: "CCC123456CCC",
			ServiceID: 19, ServiceName: "Veliki servis", ServicePrice: 35050, Available: true,
			CarServiceID: 4, CarServiceName: "Auto Servis Roki 2", Street: "Moše Pijade 17, Sremska Kamenica",
			Status: model.Rejected, SubmittedAt: time.Date(2022, 8, 15, 9, 40, 31, 0, time.Local)},

		{UserID: 3, FirstName: "Laza", LastName: "Lazic", Email: "user3@maildrop.cc",
			VehicleID: 4, Manufacturer: "Porsche", CarModel: "911 Turbo S", LicencePlate: "NS123456", ChassisNumber: "CCC123456CCC",
			ServiceID: 19, ServiceName: "Veliki servis", ServicePrice: 35050, Available: true,
			CarServiceID: 4, CarServiceName: "Auto Servis Roki 2", Street: "Moše Pijade 17, Sremska Kamenica",
			Status: model.Accepted, SubmittedAt: time.Date(2022, 8, 17, 9, 40, 31, 0, time.Local)},
	}
	appointments = []model.Appointment{
		{AppointmentRequestID: 3, WorkerID: 3, FirstName: "Pera", LastName: "Perić", StartTime: time.Date(2022, 7, 15, 10, 0, 0, 0, time.Local), FinishTime: time.Date(2022, 7, 16, 14, 34, 44, 0, time.Local), Status: model.Finished},
		{AppointmentRequestID: 4, WorkerID: 4, FirstName: "Gaša", LastName: "Gašić", StartTime: time.Date(2022, 7, 20, 10, 0, 0, 0, time.Local), FinishTime: time.Date(2022, 7, 23, 15, 27, 49, 0, time.Local), Status: model.Finished},

		{AppointmentRequestID: 7, WorkerID: 1, FirstName: "Gliša", LastName: "Glišić", StartTime: time.Date(2022, 7, 29, 12, 30, 0, 0, time.Local), FinishTime: time.Date(2022, 8, 1, 12, 30, 27, 0, time.Local), Status: model.Finished},
		{AppointmentRequestID: 8, WorkerID: 2, FirstName: "Raša", LastName: "Rašić", StartTime: time.Date(2022, 8, 25, 15, 0, 0, 0, time.Local), FinishTime: time.Time{}, Status: model.Scheduled},
		{AppointmentRequestID: 9, WorkerID: 6, FirstName: "Ivan", LastName: "Ivanovic", StartTime: time.Date(2022, 8, 16, 14, 0, 0, 0, time.Local), FinishTime: time.Time{}, Status: model.CancelledAppointment},

		{AppointmentRequestID: 11, WorkerID: 1, FirstName: "Gliša", LastName: "Glišić", StartTime: time.Date(2022, 8, 26, 9, 0, 0, 0, time.Local), FinishTime: time.Time{}, Status: model.Scheduled},
		{AppointmentRequestID: 14, WorkerID: 1, FirstName: "Gliša", LastName: "Glišić", StartTime: time.Date(2022, 8, 19, 9, 0, 0, 0, time.Local), FinishTime: time.Date(2022, 8, 21, 9, 0, 0, 0, time.Local), Status: model.Finished},
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
