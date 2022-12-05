package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type API struct{}

type BooksParams struct {
	Offset int `query:"offset"`
	Limit  int `query:"limit"`
}

type BookIdParams struct {
	ID int `param:"id"`
}

type PostBook struct {
	Title string `json:"title"`
}

var (
	books = []string{"Book 1", "Book 2", "Book 3"}
)

func (a *API) getBooks(c echo.Context) error {
	params := &BooksParams{}

	err := c.Bind(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid query params")
	}

	if params.Offset > len(books) || params.Offset < 0 {
		return c.JSON(http.StatusBadRequest, "Invalid query params")
	}

	if params.Limit < 0 || params.Limit > len(books) {
		return c.JSON(http.StatusBadRequest, "Invalid query params")
	}

	var from, to int

	if params.Offset > 0 {
		from = params.Offset
	}

	if params.Limit > 0 {
		to = params.Limit
	} else {
		to = len(books)
	}

	return c.JSON(http.StatusOK, books[from:to])
}

func (a *API) getBook(c echo.Context) error {

	params := &BookIdParams{}

	err := c.Bind(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid parameters")
	}

	index := params.ID - 1

	if index < 0 || index > len(books)-1 {
		return c.JSON(http.StatusBadRequest, "Invalid parameters")
	}

	return c.JSON(http.StatusOK, books[index])
}

func (a *API) postBook(c echo.Context) error {
	book := &PostBook{}

	err := c.Bind(book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid parameters")
	}

	books = append(books, book.Title)
	return c.NoContent(http.StatusCreated)

}
