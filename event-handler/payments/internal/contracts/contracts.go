package contracts

import "github.com/labstack/echo/v4"

type HttpHandler interface {
	RegisterRoutes(e *echo.Echo)
}
