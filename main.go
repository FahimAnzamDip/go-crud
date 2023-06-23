package main

import (
	"log"
	"net/http"

	"github.com/fahimanzamdip/go-crud/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
