package web

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

//Run web server
func Run() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/smzdm", smzdm)
	e.GET("/smzdm/list", listSmzdm)

	t := &Template{
		templates: template.Must(template.ParseGlob("web/views/*.html")),
	}
	e.Renderer = t

	e.Logger.Fatal(e.Start(":9998"))
}

//Template Template
type Template struct {
	templates *template.Template
}

//Render Render Template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
