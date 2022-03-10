package api

import "gotests/internal/books"

type API struct {
	bookService books.BookService
}

func New(service books.BookService) *API {
	return &API{
		bookService: service,
	}
}
