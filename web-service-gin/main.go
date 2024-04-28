package main

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"artist"`
	Price  float64 `json:"price"`
}

var books = []book{
	{ID: "1", Title: "Book 1", Author: "Author 1", Price: 56.99},
	{ID: "2", Title: "Book 2", Author: "Author 2", Price: 17.99},
	{ID: "3", Title: "Book 3", Author: "Author 3", Price: 39.99},
}

// getBooks responds with the list of all books as JSON.
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	index := slices.IndexFunc(books, func(b book) bool { return b.ID == id })

	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, books[index])
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postBook)
	router.Run("localhost:8080")
}
