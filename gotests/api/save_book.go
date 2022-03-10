package api

import (
	"gotests/api/dto"
	"gotests/internal/books/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *API) SaveBook(c echo.Context) error {
	ctx := c.Request().Context()
	params := dto.SaveBook{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	b := &models.Book{
		Title:  params.Title,
		Author: params.Author,
	}

	err := a.bookService.SaveBook(ctx, b)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.NoContent(http.StatusCreated)
}
