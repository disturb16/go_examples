package api

import "github.com/labstack/echo/v4"

func RegisterRoutes(a *API, e *echo.Echo) {
	e.GET("/books", a.Books)
	e.GET("/books/:id", a.BookByID)
	e.POST("/books", a.SaveBook)
}
