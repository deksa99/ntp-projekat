package database

import (
	"UserService/model"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	accounts = []model.Account{
		{Username: "admin", Password: "adminpass", Active: true, BlockedUntil: time.Time{}},
		{Username: "user1", Password: "user1pass", Active: true, BlockedUntil: time.Time{}},
		{Username: "user2", Password: "user2pass", Active: true, BlockedUntil: time.Time{}},
		{Username: "user3", Password: "user3pass", Active: false, BlockedUntil: time.Time{}},
		{Username: "worker1", Password: "worker1pass", Active: true, BlockedUntil: time.Time{}},
		{Username: "worker2", Password: "worker2pass", Active: true, BlockedUntil: time.Time{}},
	}

	admins = []model.Admin{
		{FirstName: "Dejan", LastName: "Todorovic", AccountID: 1},
	}

	users = []model.User{
		{FirstName: "Mika", LastName: "Mikic", Email: "mika@maildrop.cc", AccountID: 2},
		{FirstName: "Vera", LastName: "Veric", Email: "vera@maildrop.cc", AccountID: 3},
		{FirstName: "Laza", LastName: "Lazic", Email: "laza@maildrop.cc", AccountID: 4},
	}

	serviceWorkers = []model.ServiceWorker{
		{FirstName: "Mika", LastName: "Mikic", Main: true, AccountID: 2},
		{FirstName: "Mika", LastName: "Mikic", Main: false, AccountID: 2},
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

	err = Db.Migrator().DropTable(&model.Account{})
	if err != nil {
		log.Panic(err.Error())
	}
	err = Db.Migrator().DropTable(&model.User{})
	if err != nil {
		log.Panic(err.Error())
	}
	err = Db.Migrator().DropTable(&model.ServiceWorker{})
	if err != nil {
		log.Panic(err.Error())
	}
	err = Db.Migrator().DropTable(&model.Admin{})
	if err != nil {
		log.Panic(err.Error())
	}

	log.Println("Tables dropped successfully")

	err = Db.AutoMigrate(&model.Account{})
	if err != nil {
		log.Panic(err.Error())
	}
	err = Db.AutoMigrate(&model.User{})
	if err != nil {
		log.Panic(err.Error())
	}
	err = Db.AutoMigrate(&model.ServiceWorker{})
	if err != nil {
		log.Panic(err.Error())
	}
	err = Db.AutoMigrate(&model.Admin{})
	if err != nil {
		log.Panic(err.Error())
	}

	log.Println("Migration successful")

	for _, account := range accounts {
		Db.Create(&account)
	}

	for _, admin := range admins {
		Db.Create(&admin)
	}

	for _, user := range users {
		Db.Create(&user)
	}

	for _, serviceWorker := range serviceWorkers {
		Db.Create(&serviceWorker)
	}

	log.Println("DB seeding successful")
}
