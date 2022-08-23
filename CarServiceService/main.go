package main

import (
	"CarServiceService/router"
	"CarServiceService/util/database"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.Print("Starting CarServiceService...")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.ConnectToDatabase()
	router.HandleRequests()
	log.Print("CarServiceService started")
}
