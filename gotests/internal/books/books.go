package books

import (
	"context"
	"gotests/internal/books/models"
)

// DBIneractor is an interface for interacting with the database.
type DBInteractor interface {
	AllBooks(ctx context.Context) ([]*models.Book, error)
	BookByID(ctx context.Context, id int64) (*models.Book, error)
	SaveBook(ctx context.Context, book *models.Book) error
}

// BookService handles all the business logic for the books.
type BookService interface {
	ListBooks(ctx context.Context) ([]*models.Book, error)
	Book(ctx context.Context, id int64) (*models.Book, error)
	SaveBook(ctx context.Context, book *models.Book) error
}
