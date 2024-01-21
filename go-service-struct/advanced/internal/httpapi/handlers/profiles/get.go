package profiles

import "github.com/labstack/echo/v4"

func (h *handler) Get(c echo.Context) error {

	h.service.Get(c.Param("id"))

	return nil

}
