package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/fahimanzamdip/go-crud/models"
)

var Book models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetBooks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(books)
}
