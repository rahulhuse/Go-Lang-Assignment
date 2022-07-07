package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "The Secret", Author: "Rhonda Byrne", Quantity: 3},
	{ID: "2", Title: "The ALCHEMIST", Author: "Paulo Coelho", Quantity: 8},
	{ID: "3", Title: "Think and Grow Rich", Author: "Napoleon Hill", Quantity: 5},
}

func getAllBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)

}

func addBook(context *gin.Context) {

	var newBook book

	if err := context.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	context.IndentedJSON(http.StatusCreated, newBook)

}

func bookById(context *gin.Context) {
	id := context.Param("id")
	book, err := getBookByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, book)

}

func getBookByID(id string) (*book, error) {

	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")

}

func checkOutBook(context *gin.Context) {
	id, ok := context.GetQuery("id")

	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}

	book.Quantity -= 1
	context.IndentedJSON(http.StatusOK, book)

}

func deleteBookByID(context *gin.Context) {
	id := context.Param("id")

	_, err := getBookByID(id)

	if err != nil {

		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return

	}

	for i, b := range books {

		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
		}

	}

	context.IndentedJSON(http.StatusOK, books)

}

func main() {

	router := gin.Default()
	router.GET("/books", getAllBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", addBook)
	router.PATCH("/checkout", checkOutBook)
	router.DELETE("/deletebook/:id", deleteBookByID)
	router.Run("localhost:8080")

}
