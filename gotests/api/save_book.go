package api

import (
	"encoding/json"
	"gotests/api/dto"
	"gotests/internal/books/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *API) SaveBook(c echo.Context) error {
	ctx := c.Request().Context()
	params := dto.SaveBook{}

	data := c.Request().Body
	decoder := json.NewDecoder(data)
	err := decoder.Decode(&params)

	if err != nil {
		log.Println("Error in request", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	b := &models.Book{
		Title:  params.Title,
		Author: params.Author,
	}

	err = a.bookService.SaveBook(ctx, b)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.NoContent(http.StatusCreated)
}
