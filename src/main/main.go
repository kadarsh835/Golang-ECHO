package main

import (
	"database/dbconnect"
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var db *sql.DB
	db = dbconnect.ConnectDatabase()

	// defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}
	if err == nil {
		fmt.Println("No error for db in main!!")
	}
}
