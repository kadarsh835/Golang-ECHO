package api

import (
	"api/handlers"

	"github.com/labstack/echo"
)

// MainGroup : Export Maingroup
func MainGroup(e *echo.Echo) {
	e.GET("/", handlers.IndexHandler)
}
