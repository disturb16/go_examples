package models

// Book represents a book in the database.
type Book struct {
	ID     int64  `json:"-" db:"id"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
}
