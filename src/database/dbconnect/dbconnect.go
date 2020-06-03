package dbconnect

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	// Postgres support
	_ "github.com/lib/pq"
)

var db *sql.DB

//ConnectDatabase initializes the database
func ConnectDatabase() *sql.DB {
	databaseHost := os.Getenv("DB_HOST")
	databasePort := os.Getenv("DB_PORT")
	databaseUser := os.Getenv("DB_USER")
	databasePass := os.Getenv("DB_PASS")
	databaseName := os.Getenv("DB_NAME")

	postgresConnectionURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", databaseUser, databasePass, databaseHost, databasePort, databaseName)

	var err error
	db, err = sql.Open("postgres", postgresConnectionURL)
	if err != nil {
		log.Panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	maxOpenConn, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
	if err != nil {
		log.Panic(err)
	}
	maxIdleConn, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
	if err != nil {
		log.Panic(err)
	}

	db.SetMaxOpenConns(maxOpenConn)
	db.SetMaxIdleConns(maxIdleConn)

	log.Println("Database connected!")
	return db
}
