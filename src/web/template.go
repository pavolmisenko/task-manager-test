package web

import (
	"html/template"
	"io"
	"net/http"
	"path"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var templateRoot = "web/templates"
var templates = template.Must(template.ParseGlob(path.Join(templateRoot, "*", "*.html")))

func RenderTemplate(w http.ResponseWriter, title string, data interface{}) {
	err := templates.ExecuteTemplate(w, title, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
