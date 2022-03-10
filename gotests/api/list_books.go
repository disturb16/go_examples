package api

import (
	"gotests/api/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *API) Books(c echo.Context) error {
	ctx := c.Request().Context()
	params := dto.AllBooks{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	books, err := a.bookService.ListBooks(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, books)
}
