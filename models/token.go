package models

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
)

/*
JWT claims struct
*/
type Claims struct {
	Sub   uint
	Email string
}



func ValidateToken(token string)  (claims jwt.MapClaims, err error) {
	secret := os.Getenv("token_password")
	if secret == "" {
		secret = "secret"
	}
	claims = jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err == nil {
		log.Print("token valid")
	}
	return
}