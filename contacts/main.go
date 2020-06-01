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
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/api/v1/contacts/new", handlers.CreateContact).Methods("POST")
	r.HandleFunc("/api/v1/contacts/{userID}", handlers.GetContacts).Methods("GET")
	r.HandleFunc("/api/v1/user/new", handlers.CreateAccount).Methods("POST")
	defer models.GetDB().Close()
	log.Println("Starting server : 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!\n")
}
