package main

import (
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
	c.JSON(http.StatusOK, books)
}

func main(){
	r := gin.Default()
	r.GET("/books", getBooks)
	r.Run(":8080")
}