package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/farzamalam/gopher-exercises/contacts/models"
)

func main() {
	http.HandleFunc("/", home)
	log.Println("Starting server : 8080")
	defer models.GetDB().Close()
	contacts := models.GetContacts(1)
	fmt.Println(*contacts[0])
	fmt.Println(*contacts[1])

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!")
}
