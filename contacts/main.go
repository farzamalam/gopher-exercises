package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	log.Println("Starting server : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	defer GetDB().Close()
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!")
}
