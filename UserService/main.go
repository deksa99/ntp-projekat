package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	log.Print("Starting UserService...")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	log.Print("UserService started")
}
