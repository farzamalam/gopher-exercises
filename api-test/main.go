package main

import (
	"log"
	"net/http"

	"github.com/farzamalam/gopher-exercises/api-test/handler"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Get called!"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"Post called!"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message":"Put called!"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Delete called!"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Not found"}`))
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Get).Methods(http.MethodGet)
	r.HandleFunc("/", handler.Post).Methods(http.MethodPost)
	r.HandleFunc("/", handler.Put).Methods(http.MethodPut)
	r.HandleFunc("/", handler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/", handler.NotFound)
	log.Println("Starting server at : 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
