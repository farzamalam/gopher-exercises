package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserID int
	jwt.StandardClaims
}
