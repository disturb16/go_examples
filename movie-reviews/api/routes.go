package api

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, api *API) {
	e.POST("/users", api.Register)
}
