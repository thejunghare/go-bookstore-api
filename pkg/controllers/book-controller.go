package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thejunghare/go-bookstore/pkg/models"
)

var book []models.Book

// Display all books
func GetBooks(c *gin.Context) {
	allBooks := models.GetBooks()
	c.IndentedJSON(http.StatusOK, allBooks)
}

// Display by ID
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	// Convert string to int64
	bookID, _ := strconv.ParseInt(id, 0, 0)
	booKDetails, _ := models.GetBookByID(bookID)
	c.IndentedJSON(http.StatusOK, booKDetails)
}

// Create book
func CreateBook(c *gin.Context) {
	createBook := &models.Book{}
	if err := c.ShouldBindJSON(createBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
		return
	}

	createdBook := createBook.CreateBook()
	c.JSON(http.StatusOK, createdBook)

}

// Update
func UpdateBook(c *gin.Context){
	var updateBook = &models.Book{}
	if err := c.ShouldBindJSON(updateBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
		return
	}

	id := c.Param("id")
	bookID, _ := strconv.ParseInt(id, 0, 0)

	bookDetails, db := models.GetBookByID(bookID)
	if bookDetails.Name != ""{
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication  != ""{
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	c.IndentedJSON(http.StatusOK,bookDetails)
}

// Delete book
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	// Convert string to int64
	bookID, _ := strconv.ParseInt(id, 0, 0)
	book := models.DeleteBook(bookID)
	c.IndentedJSON(http.StatusOK, book)
}
