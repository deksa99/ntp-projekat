package database

import (
	"VehicleService/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	vehicles = []model.Vehicle{
		{UserID: 1, Manufacturer: model.Honda, CarModel: "Civic", ChassisNumber: "ABC123456ABC", LicensePlate: "NS123ABC", Color: "Blue"},
		{UserID: 1, Manufacturer: model.Chevrolet, CarModel: "Camaro", ChassisNumber: "CBA654321ABC", LicensePlate: "BG123BG", Color: "Red"},
		{UserID: 2, Manufacturer: model.Ford, CarModel: "Mustang", ChassisNumber: "AAA123456AAA", LicensePlate: "NS555CBA", Color: "Black"},
		{UserID: 3, Manufacturer: model.Porsche, CarModel: "911 Turbo S", ChassisNumber: "CCC123456CCC", LicensePlate: "NS123456", Color: "White"},
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

	err = Db.Migrator().DropTable(&model.Vehicle{})
	if err != nil {
		log.Panic(err.Error())
	}

	log.Println("Tables dropped successfully")

	err = Db.AutoMigrate(&model.Vehicle{})
	if err != nil {
		log.Panic(err.Error())
	}

	log.Println("Migration successful")

	for _, account := range vehicles {
		Db.Create(&account)
	}

	log.Println("DB seeding successful")
}
