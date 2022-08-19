package main

import (
	"ReviewService/router"
	"ReviewService/util/database"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.Print("Starting ReviewService...")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.ConnectToDatabase()
	router.HandleRequests()
	log.Print("ReviewService started")
}
