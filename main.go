package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Renderer = NewTemplate()

	e.Static("/scripts", "scripts")

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	e.GET("/add-input", func(c echo.Context) error {
		return c.Render(http.StatusOK, "input", nil)
	})

	e.POST("/submit", func(c echo.Context) error {
		var data map[string][]string
		_ = c.Bind(&data)

		js, _ := json.Marshal(data)

		return c.HTML(http.StatusOK, string(js))
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func NewTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("src/*.html")),
	}
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
