package handlers

import (
	"database/dbconnect"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// PutEventDetails : Update a particular event
func PutEventDetails(c echo.Context) error {

	// Connect to database
	var db *sql.DB
	db = dbconnect.ConnectDatabase()

	idStr, _ := strconv.Atoi(c.Param("id"))
	id := uint16(idStr)

	event := new(Event)
	if err := c.Bind(event); err != nil {
		return err
	}
	defer db.Close()
	sqlStatement := "UPDATE events SET title=$1, prize=$2, head=$3, phone=$4 WHERE id=$5"
	res, err := db.Query(sqlStatement, event.Title, event.Prize, event.Head, event.Phone, id)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, event)
	}
	return c.String(http.StatusOK, "UPDATED")
}
