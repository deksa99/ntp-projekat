package database

import (
	"ReviewService/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	reviews = []model.Review{
		{AppointmentID: 1, UserID: 1, FirstName: "Mika", LastName: "Mikic", ServiceID: 11, ServiceName: "Monteža gume", CarServiceID: 2, CarServiceName: "Mile i sin", Rating: 4, Comment: "OK", Inappropriate: false},
		{AppointmentID: 2, UserID: 1, FirstName: "Mika", LastName: "Mikic", ServiceID: 1, ServiceName: "Balansiranje", CarServiceID: 2, CarServiceName: "Mile i sin", Rating: 5, Comment: "Ekstra brzo i korektno", Inappropriate: false},
		{AppointmentID: 3, UserID: 2, FirstName: "Vera", LastName: "Veric", ServiceID: 4, ServiceName: "Zamena sijalica", CarServiceID: 1, CarServiceName: "Kvačilo", Rating: 1, Comment: "Došla sam da zamenim sijalicu, došlo je toga da majstor 'mora da zameni' nekakav zmijac, zamajac, šta ti ja znam... Katastrofa!", Inappropriate: true},
	}

	reports = []model.ReviewReport{
		{ReviewID: 1, Review: model.Review{}, Processed: true},
		{ReviewID: 3, Review: model.Review{}, Processed: true},
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

	err = Db.Migrator().DropTable(&model.Review{})
	if err != nil {
		log.Panic(err.Error())
	}
	err = Db.Migrator().DropTable(&model.ReviewReport{})
	if err != nil {
		log.Panic(err.Error())
	}

	log.Println("Tables dropped successfully")

	err = Db.AutoMigrate(&model.Review{})
	if err != nil {
		log.Panic(err.Error())
	}
	err = Db.AutoMigrate(&model.ReviewReport{})
	if err != nil {
		log.Panic(err.Error())
	}

	log.Println("Migration successful")

	for _, r := range reviews {
		Db.Create(&r)
	}

	for _, r := range reports {
		Db.Create(&r)
	}

	log.Println("DB seeding successful")
}
