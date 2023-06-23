package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fahimanzamdip/go-crud/models"
	"github.com/fahimanzamdip/go-crud/utils"
	"github.com/gorilla/mux"
)

var Book models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetBooks()
	response, err := json.Marshal(books)
	utils.CheckNilError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	ID, err := strconv.ParseInt(id, 10, 64)
	utils.CheckNilError(err)

	book, _ := models.GetBookByID(ID)
	response, err := json.Marshal(book)
	utils.CheckNilError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func StoreBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)

	newBook := book.CreateBook()
	response, err := json.Marshal(newBook)
	utils.CheckNilError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatedBook := &models.Book{}
	utils.ParseBody(r, updatedBook)

	params := mux.Vars(r)
	id := params["id"]

	ID, err := strconv.ParseInt(id, 10, 64)
	utils.CheckNilError(err)

	book, db := models.GetBookByID(ID)

	if updatedBook.Name != "" {
		book.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		book.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		book.Publication = updatedBook.Publication
	}

	db.Save(&book)

	response, err := json.Marshal(book)
	utils.CheckNilError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	ID, err := strconv.ParseInt(id, 10, 64)
	utils.CheckNilError(err)

	book := models.DeleteBook(ID)
	response, err := json.Marshal(book)
	utils.CheckNilError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}