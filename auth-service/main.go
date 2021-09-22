package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type BasicAuthResponse struct {
	Authenticated bool   `json:"authenticated"`
	User          string `json:"user"`
}

type CreateAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateAuthResponse struct {
	Created bool   `json:"created"`
	User    string `json:"user"`
}

func main() {

	http.HandleFunc("/api/v1/verify", verify)
	http.HandleFunc("/api/v1/create", create)
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

	v := verifyInDB(username, password)
	if v {
		out := BasicAuthResponse{
			Authenticated: v,
			User:          username,
		}
		resp, _ := json.Marshal(out)
		fmt.Fprint(w, string(resp))
		return
	}

}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "Invalid method\n")
		return
	}
	var b CreateAuth

	err := json.NewDecoder(r.Body).Decode(&b)
	defer r.Body.Close()

	if err != nil {
		fmt.Fprintf(w, "Invalid body: %s\n", err)
		return
	}
	log.Printf("Username: %s\n", b.Username)
	log.Printf("Password: %s\n", b.Password)
	err = createInDB(b.Username, b.Password)
	if err != nil {
		fmt.Fprintf(w, "Internal server Error: %s", err)
		return
	}
	data := CreateAuthResponse{
		Created: true,
		User:    b.Username,
	}
	resp, _ := json.Marshal(data)
	fmt.Fprintf(w, "%s\n", resp)
}

func verifyInDB(username, password string) bool {
	return false
}
func createInDB(username, password string) error {
	return nil
}
