package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-rest-api/pkg/db"
	"go-rest-api/pkg/handlers"
	"go-rest-api/pkg/utils"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Message struct {
	Text string
}

func routes() {
	log.SetOutput(os.Stdout)
	err := godotenv.Load(".env")

	db.Connect()
	if utils.CheckErr(err) {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", handlers.GetHandler).Methods(http.MethodGet)
	r.HandleFunc("/db", handlers.GetDbURL).Methods(http.MethodGet)
	r.HandleFunc("/{yyyy}/{mm}/{dd}", handlers.GetByDateHandler).Methods(http.MethodGet)
	r.HandleFunc("/{yyyy}/{mm}/{dd}", handlers.CreateByDateHandler).Methods(http.MethodPost)
	r.HandleFunc("/{yyyy}/{mm}/{dd}/entries", handlers.CreateByDateBulkHandler).Methods(http.MethodPost)
	r.HandleFunc("/", handlers.NotFound)

	http.Handle("/", r)

	PORT := os.Getenv("port")
	if PORT == "" {
		PORT = "5050"
	}
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

// HandleRequest prints message data
func HandleRequest(ctx context.Context, data Message) (string, error) {
	return fmt.Sprintf("Hello %s!", data.Text), nil
}

func main() {
	lambda.Start(HandleRequest)
}
