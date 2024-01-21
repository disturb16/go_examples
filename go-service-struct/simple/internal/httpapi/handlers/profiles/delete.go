package profiles

import "github.com/labstack/echo/v4"

func (h *handler) Delete(c echo.Context) error {

	h.service.Get(c.Param("id"))

	return c.JSON(200, "success")

}
