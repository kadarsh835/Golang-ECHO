package handlers

import (
	"database/dbconnect"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// Event : Might need to export
type Event struct {
	Id    uint16 `json :"id"`
	Title string `json :"title"`
	Prize uint16 `json :"prize"`
	Head  string `json :"head"`
	Phone string `json :"phone"`
}

// Events : Might need to export
type Events struct {
	Events []Event `json:"events"`
}

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
	rows, err := db.Query(sqlStatement)
	checkError(err)

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
	fmt.Println(result)
	db.Close()
	err = c.Render(http.StatusOK, "index.html", result)
	return err
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
