package main

import (
	"log"
	"net/http"

	"github.com/farzamalam/gopher-exercises/bookdata/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	port := "8080"
	api.HandleFunc("", handler.Home)
	api.HandleFunc("/books/isbn/{isbn}", handler.SearchByISBN).Methods(http.MethodGet)
	api.HandleFunc("/books/isbn/{isbn}", handler.Delete).Methods(http.MethodDelete)
	api.HandleFunc("/books/book/{book}", handler.SearchByBookName).Methods(http.MethodGet)
	api.HandleFunc("/books/author/{author}", handler.SearchByAuthor).Methods(http.MethodGet)
	api.HandleFunc("/books/isbn/{isbn}", handler.UpdateBook).Methods(http.MethodPut)
	api.HandleFunc("/book", handler.CreateBook).Methods(http.MethodPost)

	log.Println("Starting server at :", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
