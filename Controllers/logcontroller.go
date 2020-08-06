// controllers/books.go

package controllers

import (
	"net/http"

	models "wordplay/Models"

	"github.com/gin-gonic/gin"
)

// GET /books
// Get all books

func Listerror(c *gin.Context) {
	var logger []models.Logger
	models.DB.Find(&logger)

	c.JSON(http.StatusOK, gin.H{"data": logger})
}

// func CreateBook(c *gin.Context) {
// 	// Validate input
// 	var input models.CreateBookInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Create book
// 	book := models.Book{Title: input.Title, Author: input.Author}
// 	models.DB.Create(&book)

// 	c.JSON(http.StatusOK, gin.H{"data": book})
// }

// // GET /books/:id
// // Find a book
// func FindBook(c *gin.Context) { // Get model if exist
// 	var book models.Book

// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": book})
// }

// // PATCH /books/:id
// // Update a book
// func UpdateBook(c *gin.Context) {
// 	// Get model if exist
// 	var book models.Book
// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	// Validate input
// 	var input models.UpdateBookInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	models.DB.Model(&book).Updates(input)

// 	c.JSON(http.StatusOK, gin.H{"data": book})
// }
