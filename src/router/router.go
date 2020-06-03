package router

import (
	"api"
	"html/template"
	"io"

	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

// New : Create a new echo instance and export
func New() *echo.Echo {
	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("src/templates/crud/html/*")),
	}
	e.Renderer = renderer

	// Setup API endpoints for the main routes
	api.MainGroup(e)

	return e
}
