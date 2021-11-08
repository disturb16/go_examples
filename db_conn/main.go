package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Author struct {
	ID   int64
	Name string
}

type Book struct {
	ID          int64
	Title       string
	AuthorID    string
	PublishDate time.Time
	Author      Author
}

func main() {

	ctx := context.Background()

	db, err := createConnection()
	if err != nil {
		panic(err)
	}

	err = queryBooks(ctx, db)
	if err != nil {
		panic(err)
	}

	db.Close()
}

func createConnection() (*sql.DB, error) {
	connectionString := "root:1234@tcp(localhost:3306)/books?parseTime=True"

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func queryBooks(ctx context.Context, db *sql.DB) error {
	qry := `
		select
			b.id,
			b.title,
			b.author_id,
			a.name as 'author',
			b.publish_date
		from books b
		inner join authors a
				on a.id = b.author_id`

	rows, err := db.QueryContext(ctx, qry)
	if err != nil {
		return err
	}

	books := []Book{}

	for rows.Next() {
		b := Book{}

		err = rows.Scan(&b.ID, &b.Title, &b.AuthorID, &b.Author.Name, &b.PublishDate)
		if err != nil {
			return err
		}

		books = append(books, b)
	}

	fmt.Println(books)

	return nil
}

func addBook(ctx context.Context, db *sql.DB, title string, authorID int64, publishDate time.Time) error {

	qryadd := `insert into books(title, author_id, publish_date) values (?,?,?)`

	result, err := db.ExecContext(ctx, qryadd, title, authorID, publishDate)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	fmt.Println("INSERT ID:", id)

	return nil
}
