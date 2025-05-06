package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"size:255"`
	Author string `gorm:"size:255"`
	Price  float64
}

var DB *gorm.DB

func connectdatabase() {
	dsn := "host=localhost user=postgres password=new_password dbname=bookstore port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	DB = db

	DB.AutoMigrate(&Book{})

	fmt.Println("database is connected sussesfully!")
}

func createbook(title, author string, price float64) {
	book := Book{Title: title, Author: author, Price: price}
	result := DB.Create(&book)
	if result.Error != nil {
		fmt.Println("error inserting book", result.Error)
	} else {
		fmt.Println("Book inserted successfully!")
	}
}

func Getbooks() {
	var books []Book
	DB.Find(&books)
	fmt.Println("Boks list", books)
}

func update(id uint, newprice float64) {
	var book Book
	DB.First(&book, id)
	book.Price = newprice
	DB.Save(&book)

	fmt.Println("Book updated sussesfully")
}

func delete(id uint) {
	var book Book
	DB.Delete(&book, id)
	fmt.Println("Book deleteed Sussesfully")
}
func DeleteAllBooks() {
	result := DB.Unscoped().Delete(&Book{}) // DELETE FROM books

	if result.Error != nil {
		fmt.Println("Error deleting books:", result.Error)
	} else {
		fmt.Println("All books deleted successfully!")
	}
}

func main() {
	connectdatabase()
	//createbook("golang", "Amaresha", 68)

	//Getbooks()
	//update(2, 97.9)
	//delete(14)
	DeleteAllBooks()

}
