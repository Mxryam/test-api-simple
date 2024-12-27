package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books = []Book{
	{
		ID:     "1",
		Title:  "GO PROGRAMMING",
		Author: "John",
		Year:   "2021"},
	{
		ID:     "2",
		Title:  "Gin",
		Author: "Jane",
		Year:   "2022"},
}

func main() {
	r := gin.Default()

	r.GET("/books", getBooks)
	r.GET("/books/ : id", getBookByID)
	r.POST("/books", createBook)
	r.PATCH("/books/ :id", updateBook)
	r.DELETE("/books/:id", deleteBook)

	r.Run(":8080")
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusFound, gin.H{"message": "Book not found"})
}

func createBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook);
	err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook Book

	if err := c.BindJSON(&updatedBook);
	err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"massage":"invalid data"})
		return

	}

	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook

			c.IndentedJSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.IndentedJSON(http.StatusFound, gin.H{"message": "Book not found"})
}

func deleteBook(c *gin.Context){
	id := c.Param("id")
	initialLength := len(books)

	books = filterBooks(books, func(book Book) bool{
		return book.ID != id
	})

	if len(books) == initialLength {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message" : "book deleted"})
}

	func filterBooks(input []Book, predicate func(Book) bool) []Book {
		var result []Book
		for _, book := range input {
			if predicate(book) {
				result = append(result, book)
			}
		}
	return result
}