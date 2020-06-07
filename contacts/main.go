package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/farzamalam/gopher-exercises/contacts/app"
	"github.com/farzamalam/gopher-exercises/contacts/handlers"
	"github.com/farzamalam/gopher-exercises/contacts/models"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(app.JwtAuthentication)
	r.HandleFunc("/", home).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/contacts/new", handlers.CreateContact).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/contacts/{userID}", handlers.GetContacts).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/user/new", handlers.CreateAccount).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/user/{accountsID}", handlers.GetAccount).Methods(http.MethodGet)
	defer models.GetDB().Close()
	port := "8080"
	log.Println("Starting server : ", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!\n")
}
