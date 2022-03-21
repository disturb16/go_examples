package books

import (
	"context"
	"errors"
	"gotests/internal/books/models"
)

type bookSerive struct {
	dbInteractor DBInteractor
}

var (
	// ErrInvaliBookdID is returned when the book ID is negative.
	ErrInvaliBookdID = errors.New("invalid book id")
)

// NewService creates a new book service.
func NewService(dbInteractor DBInteractor) BookService {
	return &bookSerive{dbInteractor: dbInteractor}
}

// Books returns all books.
func (s *bookSerive) ListBooks(ctx context.Context) ([]*models.Book, error) {
	return s.dbInteractor.AllBooks(ctx)
}

// Book returns a book by its ID.
func (s *bookSerive) Book(ctx context.Context, id int64) (*models.Book, error) {
	if id <= 0 {
		return nil, ErrInvaliBookdID
	}

	return s.dbInteractor.BookByID(ctx, id)
}

func (s *bookSerive) SaveBook(ctx context.Context, book *models.Book) error {
	return s.dbInteractor.SaveBook(ctx, book)
}
