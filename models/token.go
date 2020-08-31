package models

import "github.com/dgrijalva/jwt-go"

/*
JWT claims struct
*/
type Claims struct {
	UserId   uint
	Username string
	jwt.StandardClaims
}

