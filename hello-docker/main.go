package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloWorld)
	log.Println("Starting a server : 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Docker World!")
}
