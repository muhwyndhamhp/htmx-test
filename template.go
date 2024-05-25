package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

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
