package main

import (
	"VehicleService/router"
	"VehicleService/util/database"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.Print("Starting VehicleService...")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.ConnectToDatabase()
	router.HandleRequests()
	log.Print("VehicleService started")
}
