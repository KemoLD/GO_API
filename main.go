package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quality int `json:"quality"`
}

var books = []Book{
	Book{ID: "1", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quality: 5},
	Book{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee", Quality: 4},
	Book{ID: "3", Title: "1984", Author: "George Orwell", Quality: 3},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookByID(id string) (*Book, error) {
	for i, book := range books {
		if book.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func main(){
	r := gin.Default()
	r.GET("/books", getBooks)
	r.POST("/books", createBook)
	r.GET("/books/:id", bookById)
	r.Run(":8080")
}