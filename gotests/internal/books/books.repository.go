package books

import (
	"context"
	"gotests/internal/books/models"

	"github.com/jmoiron/sqlx"
)

type bookRepository struct {
	db *sqlx.DB
}

// NewRepository creates a new book repository.
func NewRepository(db *sqlx.DB) DBInteractor {
	return &bookRepository{db: db}
}

// AllBooks returns all books.
func (r *bookRepository) AllBooks(ctx context.Context) ([]*models.Book, error) {
	books := []*models.Book{}
	err := r.db.SelectContext(ctx, &books, "SELECT * FROM book_shelf.books")
	return books, err
}

// BookByID returns a book by its ID.
func (r *bookRepository) BookByID(ctx context.Context, id int64) (*models.Book, error) {
	book := models.Book{}
	err := r.db.GetContext(ctx, &book, "SELECT * FROM book_shelf.books WHERE id=$1", id)
	return &book, err
}

func (r *bookRepository) SaveBook(ctx context.Context, book *models.Book) error {
	_, err := r.db.NamedExecContext(ctx, "INSERT INTO book_shelf.books (title, author) VALUES (:title, :author)", book)
	return err
}
