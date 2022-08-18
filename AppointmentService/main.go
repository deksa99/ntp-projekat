package main

import (
	"AppointmentService/router"
	"AppointmentService/util/database"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.Print("Starting AppointmentService...")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.ConnectToDatabase()
	router.HandleRequests()
	log.Print("AppointmentService started")
}
