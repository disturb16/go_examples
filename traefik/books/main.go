package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, []map[string]string{
			{"title": "The Hitchhiker's Guide to the Galaxy", "author": "Douglas Adams"},
			{"title": "The Restaurant at the End of the Universe", "author": "Douglas Adams"},
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
