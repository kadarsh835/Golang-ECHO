package handlers

import (
	"database/dbconnect"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// AddEventHandler : handle POST request to add event
func AddEventHandler(c echo.Context) error {

	// Connect to database
	var db *sql.DB
	db = dbconnect.ConnectDatabase()

	sqlStatement := "SELECT * FROM events ORDER BY id"
	rows, err := db.Query(sqlStatement)
	var id uint16
	for rows.Next() {
		var prize uint16
		var title, head, phone string
		err2 := rows.Scan(&id, &title, &prize, &head, &phone)
		// Exit if we get an error
		if err2 != nil {
			return err2
		}
	}

	u := new(Event)
	if err := c.Bind(u); err != nil {
		return err
	}
	u.Id = id + 1
	sqlStatement = "INSERT INTO events (id, title, prize, head, phone)VALUES ($1, $2, $3, $4, $5)"
	defer db.Close()
	_, err = db.Query(sqlStatement, u.Id, u.Title, u.Prize, u.Head, u.Phone)
	if err != nil {
		fmt.Println(err)
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
