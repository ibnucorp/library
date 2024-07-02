package main

import (
	"library/controllers"
	"library/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	db, err := db.InitDB()

	if err != nil {
		panic(err)
	}

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Selamat Datang di Perpustakaan",
		})
	})

	router.GET("/books", controllers.ShowAllBooks(db))
	router.POST("/books", controllers.InputBook(db))
	router.PUT("/books/:id", controllers.UpdateBook(db))

	router.Run(":8080")
}
