package main

import (
	"golang-app/app/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := router.SetupRouter()
	r.Run() // Default port is 8080
}
