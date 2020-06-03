package handlers

import (
	"database/dbconnect"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// IndexHandler : call from mainGroup.go
func IndexHandler(c echo.Context) error {

	// Connect to database
	var db *sql.DB
	db = dbconnect.ConnectDatabase()
	err := db.Ping()
	if err != nil {
		log.Panic(err)
	}
	if err == nil {
		log.Println("Database connected!")
	}

	sqlStatement := "SELECT * FROM events"

	defer db.Close()
	rows, err := db.Query(sqlStatement)
	CheckError(err)

	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	}
	defer rows.Close()
	result := Events{}

	for rows.Next() {
		event := Event{}
		var id, prize uint16
		var title, head, phone string
		err2 := rows.Scan(&id, &title, &prize, &head, &phone)
		event.Id = id
		event.Title = title
		event.Prize = prize
		event.Head = head
		event.Phone = phone
		// Exit if we get an error
		if err2 != nil {
			return err2
		}
		result.Events = append(result.Events, event)
	}
	err = c.Render(http.StatusOK, "index.html", result)
	return err
}
