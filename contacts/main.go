package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	port := "8000"

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home")
}
