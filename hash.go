package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func initDB() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://demo_ub3u_user:TNVfdu2b716AGyj0wFfUzeK7BKB4GLMD@dpg-cus4n3qn91rc73dhmpvg-a.oregon-postgres.render.com/demo_ub3u")
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	fmt.Println("Connected to the database successfully!")

	return db
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func RegisterUser(db *sql.DB, username, password string) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, hashedPassword)
	if err != nil {
		return err
	}

	fmt.Println("User registered successfully!")
	return nil
}

func main() {
	dbConnection := initDB()

	username := "amare"
	password := "Amaresha12"

	err := RegisterUser(dbConnection, username, password)
	if err != nil {
		fmt.Println("Error registering user:", err)
	}
}
