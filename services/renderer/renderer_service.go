package renderer

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Renderer struct {
	templates *template.Template
}

func NewRenderer() *Renderer {
	return &Renderer{
		templates: template.Must(template.New("views/*.html").Funcs(template.FuncMap{
			"itterate": func(n int) int {
				return n + 1
			},
		}).ParseGlob("views/*.html")),
	}
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
