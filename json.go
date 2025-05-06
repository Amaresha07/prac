package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Secret key for signing the token
var secretKey = []byte("my_secret_key")

// Function to generate JWT token
func generateToken(username string) (string, error) {
	// Create claims (data inside the token)
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	}

	// Create token with claims and sign it
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func main() {
	token, err := generateToken("john_doe")
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("JWT Token:", token)
}
