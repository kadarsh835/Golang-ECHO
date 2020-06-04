package handlers

import (
	"database/dbconnect"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// GetEventDetails : Get event details for a particular event
func GetEventDetails(c echo.Context) error {

	// Connect to database
	var db *sql.DB
	db = dbconnect.ConnectDatabase()

	idStr, _ := strconv.Atoi(c.Param("id"))
	id := uint16(idStr)

	event := new(Event)
	if err := c.Bind(event); err != nil {
		return err
	}
	sqlStatement := "SELECT * FROM events WHERE id = $1"
	defer db.Close()
	rows, err := db.Query(sqlStatement, id)

	var prize uint16
	var title, head, phone string
	for rows.Next() {
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
	}

	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	} else {
		fmt.Println(event)
		return c.JSON(http.StatusOK, event)
	}
	return c.String(http.StatusOK, "Fetched")
}
