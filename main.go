package main

import (
	"log"
	"net/http"
	"os"

	"go-rest-api/pkg/db"
	"go-rest-api/pkg/handlers"
	"go-rest-api/pkg/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func enableCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		log.Println("************ Headers ************ ", w)
		next.ServeHTTP(w, r)
	})
}

func main() {
	err := godotenv.Load(".env")

	db.Connect()
	if utils.CheckErr(err) {
		log.Println("Error loading .env file")
	}

	r := mux.NewRouter().StrictSlash(true)
	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/", http.StripPrefix("/", fs))
	r.HandleFunc("/api/", handlers.GetHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/db", handlers.GetDbURL).Methods(http.MethodGet)
	r.HandleFunc("/api/{yyyy}/{mm}/{dd}", handlers.GetByDateHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/{yyyy}/{mm}/{dd}", handlers.CreateByDateHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/{yyyy}/{mm}/{dd}/entries", handlers.CreateByDateBulkHandler).Methods(http.MethodPost)
	r.PathPrefix("/fe").Handler(http.FileServer(http.Dir("./build")))
	r.HandleFunc("/", handlers.NotFound)
	r.Use(enableCorsMiddleware)
	// https://github.com/gorilla/mux#serving-single-page-applications

	http.Handle("/", r)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "5050"
	}
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
