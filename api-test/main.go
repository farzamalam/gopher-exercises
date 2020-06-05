package main

import (
	"log"
	"net/http"

	"github.com/farzamalam/gopher-exercises/api-test/handler"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome Home!"))
}

func main() {
	r := mux.NewRouter()
	port := "8080"
	// Added a sub router to make provisions for multiple versions.
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", handler.Get).Methods(http.MethodGet)
	api.HandleFunc("", handler.Post).Methods(http.MethodPost)
	api.HandleFunc("", handler.Put).Methods(http.MethodPut)
	api.HandleFunc("", handler.Delete).Methods(http.MethodDelete)
	api.HandleFunc("", handler.NotFound)
	api.HandleFunc("/user/{userID}/comment/{commentID}", handler.Params).Methods(http.MethodGet)
	r.HandleFunc("/", home)
	log.Println("Starting server at : ",port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
