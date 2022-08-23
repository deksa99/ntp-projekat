package main

import (
	"EmailService/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.Println("Starting EmailService")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	router.HandleRequests()
}
