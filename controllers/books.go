package controllers

import (
	"library/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShowAllBooks(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var books []models.Book
		db.Find(&books)
		ctx.JSON(200, books)
	}
}

func InputBook(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Parse and convert
		amount, err := strconv.Atoi(ctx.PostForm("amount"))
		if err != nil {
			ctx.JSON(400, gin.H{"message": "Amount must be a number"})
			return
		}

		// Create new book
		db.AutoMigrate(models.Book{})

		// Memasukkan nilai yang diinput user ke variable newBook
		newBook := models.Book{
			Title:     ctx.PostForm("title"),
			Author:    ctx.PostForm("author"),
			Publisher: ctx.PostForm("publisher"),
			City:      ctx.PostForm("city"),
			Year:      ctx.PostForm("year"),
			Amount:    amount,
		}

		// Save data and send response
		db.Create(&newBook)
		ctx.JSON(201, newBook)
	}
}
