package main

import (
	"UserService/router"
	"UserService/util/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	log.Print("Starting UserService...")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.ConnectToDatabase()
	router.HandleRequests()
	log.Print("UserService started")
}
