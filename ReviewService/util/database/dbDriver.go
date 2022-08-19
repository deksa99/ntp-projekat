package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
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

	//err = Db.Migrator().DropTable(&model.___{})
	//if err != nil {
	//	log.Panic(err.Error())
	//}
	//
	//log.Println("Tables dropped successfully")
	//
	//err = Db.AutoMigrate(&model.___{})
	//if err != nil {
	//	log.Panic(err.Error())
	//}
	//
	//log.Println("Migration successful")
	//
	//for _, account := range ___ {
	//	Db.Create(&account)
	//}
	//
	//log.Println("DB seeding successful")
}
