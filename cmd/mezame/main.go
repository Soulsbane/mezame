package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Location struct {
	Name string `json:"name"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Location{Name: "Tokyo"})
		// return c.JSONPretty(http.StatusOK, Location{Name: "Tokyo"}, "  ")
	})
	e.Logger.Fatal(e.Start(":1323"))

}
