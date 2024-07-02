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

func UpdateBook(db *gorm.DB) gin.HandlerFunc {
    return func (ctx *gin.Context){
        		ID := ctx.Param("id")
		var book models.Book

		if err := db.First(&book, "id=?", ID).Error; err != nil {
			ctx.JSON(400, gin.H{"message": "Product ID not found"}) // 400 : invalid syntax request to server
		}

		amount, err := strconv.Atoi(ctx.PostForm("amount"))
		if err != nil {
			ctx.JSON(400, gin.H{"message": "Amount must be a number"})
			return
		}

		var updatedBook = models.Book{
			ID:    book.ID,
			Title:     ctx.PostForm("title"),
			Author:    ctx.PostForm("author"),
			Publisher: ctx.PostForm("publisher"),
			City:      ctx.PostForm("city"),
			Year:      ctx.PostForm("year"),
			Amount:    amount,
		}

		db.Model(&book).Updates(updatedBook)
		ctx.JSON(200, book)
    }
}