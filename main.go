package main

import (
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}


//تعریف متغیر يایگاه داده

var db *gorm.DB

func main() {
	//اتصال به دیتابیس

	var err error
	db, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	iferr != nil {
		log.Fatal("faild to connect to database:", err)
	}
	 //مهاجرت دیتابیس
	 if err := db.AutoMigrate(&Book{});
	 err != nil {
		log.Fatal("faild to migrete database:", err)
	 }

	 //داده نمونه

	 initSampleData()

	r := gin.Default()

	r.GET("/books", getBooks)
	r.GET("/books/ : id", getBookByID)
	r.POST("/books", createBook)
	r.PATCH("/books/ :id", updateBook)
	r.DELETE("/books/:id", deleteBook)

	r.Run(":8080")
}

func initSampleData() {
	sampleBooks := []Book{
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
	for _, book := range sampleBooks {
		var existing Book
		if err := db.First(&existing, "id=?", book.ID).Error; err != nil {
			db.Create(&book)
		}
	}
}

func getBooks(c *gin.Context) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message" : "failed to retrieve books"})
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	var book Book

if err := db.First(&existing, "id=?", book.ID).Error; err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message" : "book not found"})
		return

}
c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook);
	err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}
	if err := db.Create(&newBook).Error;
	err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "faild to creat book"})
		return
	}
	
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
	var book Book
	if err := db.First(&book, "id = ?" , id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : :"book not found"})
		return
	}

	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Year = updatedBook.Year

	if err := db.Save(&book).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message" : "failed to update book"})
	}
	c.IndentedJSON(http.StatusOK, book)

}

func deleteBook(c *gin.Context){
	id := c.Param("id")
	if err := db.Delete(&book{}, "id = ?" , id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "failed to delete book"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message" : "book deleted"})
	
}

	