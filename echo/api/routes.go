package api

import (
	"github.com/labstack/echo/v4"
)

func (a *API) RegisterRoutes(e *echo.Echo) {

	e.Use(requestIDHandler)

	public := e.Group("")
	protected := e.Group("", authMiddleware)

	public.GET("/books", a.getBooks)
	public.GET("/books/:id", a.getBook)

	protected.POST("/books", a.postBook)
}
