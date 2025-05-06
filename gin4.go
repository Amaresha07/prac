package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database instance
var DB *gorm.DB

// Book Model
type Book struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"size:255"`
	Author string `gorm:"size:255"`
	Price  float64
}

// Connect to database
func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=new_password dbname=bookstore port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	DB.AutoMigrate(&Book{}) // Create table

	fmt.Println("Database connected and migrated!")
}

func main() {
	ConnectDatabase()

	r := gin.Default()

	// Create a new book
	r.POST("/books", func(c *gin.Context) {
		var book Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		DB.Create(&book)
		c.JSON(http.StatusOK, gin.H{"message": "Book added!", "book": book})
	})

	// Get all books
	r.GET("/books", func(c *gin.Context) {
		var books []Book
		DB.Find(&books)
		c.JSON(http.StatusOK, books)
	})

	// Update a book
	r.PUT("/books/:id", func(c *gin.Context) {
		var book Book
		id := c.Param("id")

		if err := DB.First(&book, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		DB.Save(&book)
		c.JSON(http.StatusOK, gin.H{"message": "Book updated!", "book": book})
	})

	// Delete a book
	r.DELETE("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		DB.Delete(&Book{}, id)
		c.JSON(http.StatusOK, gin.H{"message": "Book deleted!"})
	})

	r.Run(":8080")
}
