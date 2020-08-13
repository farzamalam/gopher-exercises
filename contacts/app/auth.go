package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/farzamalam/gopher-exercises/contacts/models"
	"github.com/farzamalam/gopher-exercises/contacts/utils"
)

func JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/api/v1/user/new", "/api/v1/user/login", "/api/v1/user"} // List of end points	that doesn' require auth.
		log.Println(r.URL.Path)
		requestPath := r.URL.Path // Current Request path
		// Check if the request doesn't need authentication, serve the request
		for _, value := range notAuth {
			if strings.Contains(requestPath, value) {
				next.ServeHTTP(w, r)
				return
			}
		}
		tokenHeader := r.Header.Get("Authorization") // Grab the token from the request header
		// Token is missing, return with error code 403.
		if tokenHeader == "" {
			utils.Respond(w, http.StatusForbidden, utils.Message(false, "Missing Auth token."))
			return
		}
		// Token noramally comes in format `Bearer {token-body}`, Check if its in
		// Correct format.
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			utils.Respond(w, http.StatusForbidden, utils.Message(false, "Invalid or Malformed auth token."))
			return
		}
		// Grab the token part.
		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})
		if err != nil {
			utils.Respond(w, http.StatusForbidden, utils.Message(false, "Malformed authentication token."))
			return
		}
		// Token is invalid, maybe not signed on this server.
		if !token.Valid {
			utils.Respond(w, http.StatusForbidden, utils.Message(false, "Token is not valid."))
			return
		}
		// Everything went well, ServerHTTP.
		fmt.Sprintf("User %s", tk.UserID)
		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
