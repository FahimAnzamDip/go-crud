package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fahimanzamdip/go-crud/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routes.RegisterBookRoutes(router)
	http.Handle("/", router)

	fmt.Println("API URL: http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", router))
}
