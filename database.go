package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func createTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age INT NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("❌ Table creation failed:", err)
	}
	fmt.Println("✅ Table created successfully!")
}
func main() {
	createTable(db)
	connect := "host=localhost port=5432 user=myuser dbname=mydb password=mypassword"

	db, err := sql.Open("postgres", connect)
	if err != nil {
		log.Fatal("Failed to connect", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("❌ Database not reachable:", err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")
}
