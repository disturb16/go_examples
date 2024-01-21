package handlers

import "github.com/labstack/echo/v4"

type Handler interface {
	RegisterRoutes(e *echo.Echo)
}
