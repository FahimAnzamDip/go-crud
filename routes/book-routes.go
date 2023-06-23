package routes

import (
	"github.com/fahimanzamdip/go-crud/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router) {
	// router.HandleFunc("/book", controllers.StoreBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	// router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	// router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	// router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
