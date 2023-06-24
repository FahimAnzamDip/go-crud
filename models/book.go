package models

import (
	"github.com/fahimanzamdip/go-crud/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
    Name        string `gorm:"column:name;not null" json:"name"`
    Author      string `gorm:"column:author;not null" json:"author"`
    Publication string `gorm:"column:publication;not null" json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book {
	db.NewRecord(book)
	db.Create(&book)

	return book
}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)

	return Books
}

func GetBookByID(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)

	return &book, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Unscoped().Delete(&book)

	return book
}
