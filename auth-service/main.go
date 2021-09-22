package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/api/v1/verify", verify)
	port := "8080"
	log.Printf("Starting server at : %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

func verify(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		log.Printf("Empty credentials\n")
		fmt.Fprintf(w, "Empty credentials\n")
		return
	}
	log.Printf("Username: %s\n", username)
	log.Printf("Password: %s\n", password)
}
