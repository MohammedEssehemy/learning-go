package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Book struct {
	ID     int64
	Title  string
	Author string
	Price  float32
}

var db *sql.DB

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	// Get a database handle.
	var err error
	db, err = sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	books, err := booksByAuthor("Author 1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Books found: %v\n", books)

	// Hard-code ID 2 here to test the query.
	book, err := bookByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Book found: %v\n", book)

	bookID, err := addBook(Book{
		Title:  "Book 5",
		Author: "Author 4",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added book: %v\n", bookID)

}

// booksByAuthor queries for books that have the specified artist name.
func booksByAuthor(name string) ([]Book, error) {
	// A books slice to hold data from returned rows.
	var books []Book

	rows, err := db.Query("SELECT * FROM book WHERE author = $1", name)
	if err != nil {
		return nil, fmt.Errorf("booksByAuthor %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price); err != nil {
			return nil, fmt.Errorf("booksByAuthor %q: %v", name, err)
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("booksByAuthor %q: %v", name, err)
	}
	return books, nil
}

// bookByID queries for the book with the specified ID.
func bookByID(id int64) (Book, error) {
	// A book to hold data from the returned row.
	var book Book

	row := db.QueryRow("SELECT * FROM book WHERE id = $1", id)
	if err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Price); err != nil {
		if err == sql.ErrNoRows {
			return book, fmt.Errorf("bookByID %d: no such book", id)
		}
		return book, fmt.Errorf("bookByID %d: %v", id, err)
	}
	return book, nil
}

// addBook adds the specified book to the database,
// returning the book ID of the new entry
func addBook(book Book) (int64, error) {
	var id int64
	err := db.QueryRow("INSERT INTO book (title, author, price) VALUES ($1, $2, $3) RETURNING id", book.Title, book.Author, book.Price).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("addBook: %v", err)
	}
	if err != nil {
		return 0, fmt.Errorf("addBook: %v", err)
	}
	return id, nil
}
