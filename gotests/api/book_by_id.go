package api

import (
	"gotests/api/dto"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *API) BookByID(c echo.Context) error {
	ctx := c.Request().Context()
	params := dto.BookByID{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	book, err := a.bookService.Book(ctx, params.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, book)
}
