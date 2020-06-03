package main

import (
	"log"
	"router"

	"github.com/joho/godotenv"
)

func main() {
	// Load From .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect Server
	e := router.New()
	e.Start(":8080")
}
