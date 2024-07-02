package controllers

import (
	"library/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShowAllBooks(db *gorm.DB) gin.HandlerFunc {
    return func(ctx *gin.Context){
        var books []models.Book
        db.Find(&books)
        ctx.JSON(200, books)
    }
}

func InputBook(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Parse and convert stock
		amount, err := strconv.Atoi(ctx.PostForm("amount"))
		if err != nil {
			ctx.JSON(400, gin.H{"message": "Stock must be a number"})
			return
		}

		// Create new book
		var newBook models.Book
		if err := ctx.ShouldBind(&newBook); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid request body"})
			return
		}
		newBook.Amount = amount

		// Save data and send response
		db.Create(&newBook)
		ctx.JSON(201, newBook)
	}
}