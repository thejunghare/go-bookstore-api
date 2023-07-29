package models

import (
	"github.com/thejunghare/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// In order to make change to database

// GetAllBooks
func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// GetBookByID
func GetBookByID(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

// CreateBook
func (b *Book) CreateBook() *Book {
	// db.NewRecord(b)
	db.Create(&b)
	return b
}

// DeleteBook
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
