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
		{Username: "admin", Password: "$2a$12$jhJ7spSwiuuY2qxI0mAwgOiM//GFaUf9ftY5oIXo5wIAPZkSsLdYS", Active: true, BlockedUntil: time.Time{}},
		{Username: "user1", Password: "$2a$12$U0n6l/2jYCzDHkqdwnEGdOv.DdM0rWNwECIEDluN6YGeCkX6B1w62", Active: true, BlockedUntil: time.Time{}},
		{Username: "user2", Password: "$2a$12$wSZ.iF1w5UaQGnBJtcezEeQYOXiW2tNQ113nojrR1rirejbbI/lOW", Active: true, BlockedUntil: time.Time{}},
		{Username: "user3", Password: "$2a$12$AUv8/BqOs24eFcS7tka8GOoTdVi6EmOHU/Rk.fKCcd7BA5zBsQ8Lm", Active: true, BlockedUntil: time.Time{}},
		{Username: "worker1", Password: "$2a$12$gIuAQPRYpOUCCJCIsiVE6uvxd6z3eK5RHRzOZP/VfAK8v/PDY/oV.", Active: true, BlockedUntil: time.Time{}},
		{Username: "worker2", Password: "$2a$12$UwnKoVSME0ti9O9UKwTdoeu/LJs/Tl8CpV3IWK9CaM8p6KHS5JMLa", Active: true, BlockedUntil: time.Time{}},
		{Username: "worker3", Password: "$2a$12$Q8ze/7EkDUgC44OxlPj0hO2aO1lUomBfztuM00PSL9qnaKTZSTLq6", Active: true, BlockedUntil: time.Time{}},
		{Username: "worker4", Password: "$2a$12$evdNgUWNqqzBFTwVWzsPUujN7e//YXeoYuA4TQNMwR/mfMRy1g.cu", Active: true, BlockedUntil: time.Time{}},
		{Username: "worker5", Password: "$2a$12$l5iShYcENWsTV6riEAB0kOwwLlIvZiff0QOdKiVJ115M7mJfV.Oh6", Active: true, BlockedUntil: time.Time{}},
		{Username: "worker6", Password: "$2a$12$viN..5R59q000.v4sjHpj.KhoqCp.u7jDagGqNJpktFrre5FyvwKa", Active: true, BlockedUntil: time.Time{}},
	}

	admins = []model.Admin{
		{FirstName: "Dejan", LastName: "Todorovic", AccountID: 1},
	}

	users = []model.User{
		{FirstName: "Mika", LastName: "Mikic", Email: "user1@maildrop.cc", AccountID: 2},
		{FirstName: "Vera", LastName: "Veric", Email: "user2@maildrop.cc", AccountID: 3},
		{FirstName: "Laza", LastName: "Lazic", Email: "user3@maildrop.cc", AccountID: 4},
	}

	serviceWorkers = []model.ServiceWorker{
		{FirstName: "Gliša", LastName: "Glišić", Email: "worker1@maildrop.cc", Main: true, AccountID: 5, CarServiceID: 1},
		{FirstName: "Raša", LastName: "Rašić", Email: "worker2@maildrop.cc", Main: false, AccountID: 6, CarServiceID: 1},
		{FirstName: "Pera", LastName: "Perić", Email: "worker3@maildrop.cc", Main: true, AccountID: 7, CarServiceID: 2},
		{FirstName: "Gaša", LastName: "Gašić", Email: "worker4@maildrop.cc", Main: false, AccountID: 8, CarServiceID: 2},
		{FirstName: "Marko", LastName: "Markovic", Email: "worker5@maildrop.cc", Main: true, AccountID: 9, CarServiceID: 3},
		{FirstName: "Ivan", LastName: "Ivanovic", Email: "worker6@maildrop.cc", Main: true, AccountID: 10, CarServiceID: 4},
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
