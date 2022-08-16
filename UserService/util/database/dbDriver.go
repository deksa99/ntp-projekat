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
		{Username: "admin", Password: "$2a$12$mY35hEYpanAgN6NQY9m69O/arJI3.pcH9pvTZkvfv/xFL.O4Ueyse",
			Active: true, BlockedUntil: time.Time{}},
		{Username: "user1", Password: "$2a$12$U0n6l/2jYCzDHkqdwnEGdOv.DdM0rWNwECIEDluN6YGeCkX6B1w62",
			Active: true, BlockedUntil: time.Time{}},
		{Username: "user2", Password: "$2a$12$wSZ.iF1w5UaQGnBJtcezEeQYOXiW2tNQ113nojrR1rirejbbI/lOW",
			Active: true, BlockedUntil: time.Time{}},
		{Username: "user3", Password: "$2a$12$AUv8/BqOs24eFcS7tka8GOoTdVi6EmOHU/Rk.fKCcd7BA5zBsQ8Lm",
			Active: false, BlockedUntil: time.Time{}},
		{Username: "worker1", Password: "$2a$12$gIuAQPRYpOUCCJCIsiVE6uvxd6z3eK5RHRzOZP/VfAK8v/PDY/oV.",
			Active: true, BlockedUntil: time.Time{}},
		{Username: "worker2", Password: "$2a$12$UwnKoVSME0ti9O9UKwTdoeu/LJs/Tl8CpV3IWK9CaM8p6KHS5JMLa",
			Active: true, BlockedUntil: time.Time{}},
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
		{FirstName: "Mika", LastName: "Mikic", Main: true, AccountID: 5},
		{FirstName: "Mika", LastName: "Mikic", Main: false, AccountID: 6},
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
