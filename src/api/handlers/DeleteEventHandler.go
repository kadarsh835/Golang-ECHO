package handlers

import (
	"database/dbconnect"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// DeleteEventHandler : Delete an event
func DeleteEventHandler(c echo.Context) error {

	// Connect to database
	var db *sql.DB
	db = dbconnect.ConnectDatabase()

	id := c.Param("id")
	sqlStatement := "DELETE FROM events WHERE id = $1"
	defer db.Close()
	res, err := db.Query(sqlStatement, id)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	} else {
		fmt.Println(res)
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
