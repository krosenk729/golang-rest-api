package main

import (
	"log"
	"net/http"

	"go-rest-api/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", handlers.GetHandler).Methods("GET")
	r.HandleFunc("/", handlers.PostHandler).Methods("POST")
	r.HandleFunc("/{yyyy}/{mm}/{dd}", handlers.GetByDateHandler).Methods(http.MethodGet)
	r.HandleFunc("/", handlers.NotFound)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":5050", nil))
}
