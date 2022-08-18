package database

import (
	"CarServiceService/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	locations = []model.Location{
		{Longitude: 19.8298413, Latitude: 45.2618498, Street: "Bulevar oslobođenja 21, Novi Sad"},
		{Longitude: 19.8615368, Latitude: 45.2542202, Street: "Vladimira Nazora 2, Petrovaradin"},
		{Longitude: 19.8480734, Latitude: 45.2421371, Street: "Veljka Petrovica 5, Novi Sad"},
		{Longitude: 19.848706, Latitude: 45.2224511, Street: "Moše Pijade 17, Sremska Kamenica"},
	}
	carServices = []model.CarService{
		{Name: "Kvačilo", LocationID: 1},
		{Name: "Mile i sin", LocationID: 2},
		{Name: "Auto Servis Roki", LocationID: 3},
		{Name: "Auto Servis Roki 2", LocationID: 4},
	}
	services = []model.Service{
		{Name: "Mali servis", Description: "Mali servis, velika briga. Prepustite je nama.", Price: 7300, ExpectedTime: 3, CarServiceID: 1},
		{Name: "Veliki servis", Description: "Veliki servis Vašeg vozila", Price: 27600, ExpectedTime: 5, CarServiceID: 1},
		{Name: "Poliranje farova", Description: "Da blista!", Price: 1800, ExpectedTime: 1, CarServiceID: 1},
		{Name: "Zamena sijalica", Description: "LED i Xenon", Price: 850, CarServiceID: 1},
		{Name: "Zamena akulumatora", Description: "Akumulator nije uključen u cenu!", Price: 1350, CarServiceID: 1},

		{Name: "Mali servis", Description: "Mali servis vozila, cena može da varira.", Price: 8370, ExpectedTime: 2, CarServiceID: 2},
		{Name: "Veliki servis", Description: "Veliki servis vozila, cena može da varira.", Price: 29010, ExpectedTime: 4, CarServiceID: 2},
		{Name: "Mašinsko ispravljanje felni", Description: "Ko strela!", Price: 1970, ExpectedTime: 1, CarServiceID: 2},
		{Name: "Zavarivanje felne", Description: "Bez šavova", Price: 1830, CarServiceID: 2},
		{Name: "Skidanje točka", Description: "Za felne 15-17", Price: 320, CarServiceID: 2},
		{Name: "Montaža gume", Description: "Za felne 15-17", Price: 620, CarServiceID: 2},
		{Name: "Balansiranje", Description: "Za felne 15-17", Price: 880, CarServiceID: 2},

		{Name: "Mali servis", Description: "Mali servis vozila, cena može da varira.", Price: 9700, ExpectedTime: 2, CarServiceID: 3},
		{Name: "Veliki servis", Description: "Veliki servis vozila, cena može da varira.", Price: 35050, ExpectedTime: 4, CarServiceID: 3},
		{Name: "Skidanje točka", Description: "Za felne 19-20", Price: 380, CarServiceID: 3},
		{Name: "Montaža gume", Description: "Za felne 19-20", Price: 750, CarServiceID: 3},
		{Name: "Balansiranje", Description: "Za felne 19-20", Price: 990, CarServiceID: 3},

		{Name: "Mali servis", Description: "Mali servis vozila, cena može da varira.", Price: 9700, ExpectedTime: 2, CarServiceID: 4},
		{Name: "Veliki servis", Description: "Veliki servis vozila, cena može da varira.", Price: 35050, ExpectedTime: 4, CarServiceID: 4},
		{Name: "Kontrola kočnica", Description: "Kontrola kočnica po najsavremenijim standardima", Price: 1600, ExpectedTime: 1, CarServiceID: 4},
		{Name: "Kontrola amortizera", Description: "Kontrola amortizera po najsavremenijim standardima", Price: 1600, CarServiceID: 4},
		{Name: "Skidanje točka", Description: "Za felne 21-28", Price: 430, CarServiceID: 4},
		{Name: "Montaža gume", Description: "Za felne 21-28", Price: 880, CarServiceID: 4},
		{Name: "Balansiranje", Description: "Za felne 21-28", Price: 1430, CarServiceID: 4},
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

	err = Db.Migrator().DropTable(&model.Location{})
	if err != nil {
		log.Panic(err.Error())
	}
	err = Db.Migrator().DropTable(&model.CarService{})
	if err != nil {
		log.Panic(err.Error())
	}
	err = Db.Migrator().DropTable(&model.Service{})
	if err != nil {
		log.Panic(err.Error())
	}

	log.Println("Tables dropped successfully")

	err = Db.AutoMigrate(&model.Location{})
	if err != nil {
		log.Panic(err.Error())
	}
	err = Db.AutoMigrate(&model.CarService{})
	if err != nil {
		log.Panic(err.Error())
	}
	err = Db.AutoMigrate(&model.Service{})
	if err != nil {
		log.Panic(err.Error())
	}

	log.Println("Migration successful")

	for _, account := range locations {
		Db.Create(&account)
	}
	for _, account := range carServices {
		Db.Create(&account)
	}
	for _, account := range services {
		Db.Create(&account)
	}

	log.Println("DB seeding successful")
}
