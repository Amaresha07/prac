package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid data"})
			return
		}
		c.JSON(200, gin.H{"name": user.Name, "email": user.Email})
	})
	r.Run(":8080")

}
