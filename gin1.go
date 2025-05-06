package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// GET request
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{"user_id": id})
	})

	// POST request
	r.POST("/post", func(c *gin.Context) {
		c.JSON(201, gin.H{"message": "Data received!"})
	})

	// PUT request
	r.PUT("/update", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Data updated!"})
	})

	// DELETE request
	r.DELETE("/delete", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Data deleted!"})
	})

	r.Run(":8080") // Start the server
}
