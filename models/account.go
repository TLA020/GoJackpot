package models

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

// GORM defined a gorm.Model struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token" sql:"-"`
}

func (account *Account) Validate() (err error) {

	if !strings.Contains(account.Email, "@") {
		err = fmt.Errorf("email address is required")
		return
	}

	if len(account.Password) < 6 {
		err = fmt.Errorf("required and at least 6 chars")
		return
	}

	var occurrences = 0
	// check if mail unique
	err = GetDB().Table("accounts").Where("email = ?", account.Email).Count(&occurrences).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		err = fmt.Errorf("connection error")
		return
	}

	if occurrences > 0 {
		err = fmt.Errorf("address already in use by another user")
		return
	}

	return
}

func (account *Account) Create() (err error) {
	err = account.Validate()
	if err != nil {
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		err = fmt.Errorf("could not create account")
		return
	}

	//Create new JWT token for the newly registered account
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	// don't expose password
	account.Password = ""

	return
}

func (account *Account) Login() (err error) {
	var dbAccount = &Account{}

	err = GetDB().Table("accounts").Where("email = ?", account.Email).First(dbAccount).Error
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbAccount.Password), []byte(account.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		err = fmt.Errorf("error password miss match")
	}

	//Create JWT token
	token := &Token{UserId: dbAccount.ID}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), token)
	signedToken, _ := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	// don't expose password
	account.Password = ""
	account.Token = signedToken
	account.CreatedAt = dbAccount.CreatedAt
	account.UpdatedAt = dbAccount.UpdatedAt
	account.ID = dbAccount.ID
	return
}
