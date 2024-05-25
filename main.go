package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Renderer = NewTemplate()

	e.Static("/scripts", "scripts")

	e.GET("/", func(c echo.Context) error {
		res := map[string]interface{}{
			"Form": []interface{}{
				map[string]string{
					"Name": "title",
				},
				map[string]string{
					"Name": "description",
				},
			},
		}
		return c.Render(http.StatusOK, "index", res)
	})

	e.GET("/add-input", func(c echo.Context) error {
		name := c.QueryParam("name")
		res := map[string]string{"Name": name}
		return c.Render(http.StatusOK, "input", res)
	})

	e.POST("/submit", func(c echo.Context) error {
		var data map[string][]string
		_ = c.Bind(&data)

		js, _ := json.MarshalIndent(data, "", "  ")

		return c.HTML(http.StatusOK, string(js))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
