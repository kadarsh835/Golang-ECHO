package api

import (
	"api/handlers"

	"github.com/labstack/echo"
)

// MainGroup : Export Maingroup
func MainGroup(e *echo.Echo) {
	e.GET("/", handlers.IndexHandler)
	e.POST("/add_event", handlers.AddEventHandler)
	e.GET("/delete/:id", handlers.DeleteEventHandler)
	e.GET("/get_details", handlers.GetEventDetails)
	e.POST("/update_detail", handlers.PutEventDetails)
}
