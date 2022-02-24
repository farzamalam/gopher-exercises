package main

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
)

type TokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
	Scope       string `json:"scope"`
}

func main() {
	port := "8080"
	http.HandleFunc("/auth/realms/oauth/token", generateTokenHandler)
	log.Println("Starting server at: ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func generateTokenHandler(w http.ResponseWriter, r *http.Request) {
	var tokenRequest TokenRequest
	err := json.NewDecoder(r.Body).Decode(&tokenRequest)
	if err != nil {
		log.Println("error in decoding: ", err)
		return
	}
	log.Printf("Request: %+v", tokenRequest)
	tokenResponse := TokenResponse{
		AccessToken: generateRandomString(20),
		ExpiresIn:   3600,
		TokenType:   "Bearer",
		Scope:       tokenRequest.Scope,
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenResponse)
}

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomString(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return ""
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret)
}
