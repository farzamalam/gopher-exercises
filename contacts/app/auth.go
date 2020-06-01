package app

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/farzamalam/gopher-exercises/contacts/utils"
	"net/http"
)

func JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func ( w http.ResponseWriter, r *http.Request)){
		notAuth := []string{"/api/v1/user/new", "/api/v1/user/login"} // List of end points	that doesn' require auth.
		requestPath := r.URL.Path // Current Request path
		// Check if the request doesn't need authentication, serve the request
		for _, value := range notAuth{
			if value == requestPath{
				next.ServeHTTP(w, r)
				return
			}
		}
		tokenHeader := r.Header.Get("Authorization") // Grab the token from the request header
		// Token is missing, return with error code 403.
		if tokenHeader == ""{
			utils.Respond(w, http.StatusForbidden,utils.Message(false, "Missing Auth token."))
			return
		}
		// Token noramally comes in format `Bearer {token-body}`, Check if its in
		// Correct format.
		splitted := strings.Split(tokenHeader," ")
		if len(splitted) != 2{
			utils.Respond(w, http.StatusForbidden, utils.Message(false,"Invalid or Malformed auth token."))
			return
		}
		// Grab the token part.
		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart,tk,func(token *jwt.Token)(interface{}, error){
			return []byte(os.Getenv("token_password")),nil	
		})
		if err != nil{
			utils.Respond(w, http.StatusForbidden, utils.Message(false, "Malformed authentication token."))
			return
		}
		// Token is invalid, maybe not signed on this server.
		if !token.Valid {
			utils.Respond(w, http.StatusForbidden, utils.Message(false, "Token is not valid."))
			return
		}
		// Everything went well, ServerHTTP.
		fmt.Sprintf("User %s",tk.Username)
		ctx := context.WithValue(r.Context,"user",tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w,r
		)}
}
