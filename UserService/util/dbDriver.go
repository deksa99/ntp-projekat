package util

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

func ConnectToDatabase() {
	connStr := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully connected to DB")
	}

	db.Migrator().DropTable(&model.Account{})
	db.Migrator().DropTable(&model.User{})
	db.Migrator().DropTable(&model.ServiceWorker{})
	db.Migrator().DropTable(&model.Admin{})

	log.Println("Tables dropped successfully")

	db.AutoMigrate(&model.Account{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.ServiceWorker{})
	db.AutoMigrate(&model.Admin{})

	log.Println("Migration successful")

	for _, account := range accounts {
		db.Create(&account)
	}

	for _, admin := range admins {
		db.Create(&admin)
	}

	for _, user := range users {
		db.Create(&user)
	}

	for _, serviceWorker := range serviceWorkers {
		db.Create(&serviceWorker)
	}

	log.Println("DB seeding successful")
}
