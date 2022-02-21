package main

import (
	"net/http"

	"github.com/disturb16/go_examples/echo/api"
	"github.com/labstack/echo/v4"
)

func handleIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello World"})
}

func main() {
	e := echo.New()
	e.GET("/", handleIndex)

	a := &api.API{}
	a.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8081"))
}
