package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
	m "goprac/models"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var signUp = func(c *fiber.Ctx) {
	account := &m.Account{}

	if err := c.BodyParser(account); err != nil {
		log.Fatal(err)
	}

	if err := account.Validate(); err != nil {
		log.Print(err)
		c.SendStatus(fiber.StatusBadRequest)
		c.Send(err)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	m.GetDB().Create(account)

	if account.ID <= 0 {
		c.SendStatus(fiber.StatusBadRequest)
		return
	}

	//Create new JWT token for the newly registered account
	token, err := createTokenByAccount(account)
	if err != nil {
		log.Print(err)
		c.SendStatus(fiber.StatusInternalServerError)
		return
	}

	// don't expose password
	account.Password = "SECRET"
	account.Token = token

	if err := c.JSON(account); err != nil {
		log.Print(err)
	}
}

var signIn = func(c *fiber.Ctx) {
	account := &m.Account{}
	dbAccount := &m.Account{}

	if err := c.BodyParser(account); err != nil {
		log.Print(err)
		c.SendStatus(fiber.StatusBadRequest)
		return
	}
	// lookup account in db
	if err := m.GetDB().Table("accounts").Where("email = ?", account.Email).First(dbAccount).Error; err != nil {
		log.Print(err)
		c.SendStatus(fiber.StatusUnauthorized)
		c.Send("looks like this account doesn't exist")
		return
	}

	// compare encrypted passwords
	if err := bcrypt.CompareHashAndPassword([]byte(dbAccount.Password), []byte(account.Password)); err != nil {
		log.Print("passwords don't match")
		c.SendStatus(fiber.StatusUnauthorized)
		c.Send("credentials invalid")
		return
	}

	token, err := createTokenByAccount(dbAccount)
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return
	}

	dbAccount.Token = token
	dbAccount.Password = "SECRET"
	if err := c.JSON(dbAccount); err != nil {
		log.Print(err)
	}
}

func createTokenByAccount(acc *m.Account) (signedToken string, err error) {
	// Generate/Sign encoded JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = acc.Email
	claims["username"] = acc.Username
	claims["avatar"] = acc.Avatar
	claims["sub"] = acc.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	// Generate encoded token (sign)
	signedToken, err = token.SignedString([]byte(secret))
	return
}

var uploadAvatar = func(ctx *fiber.Ctx) {
	const avatarPath = "uploads/avatars"

	// get user uploading avatar from claims
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	userId := claims["sub"].(float64)

	multipartFileHeader, _ := ctx.FormFile("avatar")
	allowed := []string{"image/jpeg", "image/jpg", "image/png"}

	// validate by allowed mimetypes
	if !validateFile(multipartFileHeader, allowed) {
		log.Print("filetype not allowed")
		ctx.Status(http.StatusBadRequest).Send("invalid file type")
		return
	}

	// this ensures folder exist.
	if _, err := os.Stat(avatarPath); os.IsNotExist(err) {
		_ = os.MkdirAll(avatarPath, os.ModePerm)
	}

	// create / overwrite avatar
	extension := filepath.Ext(multipartFileHeader.Filename)
	fullPath := fmt.Sprintf("%s/%s%s", avatarPath, username, extension)
	if err := ctx.SaveFile(multipartFileHeader, fullPath); err != nil {
		log.Print(err)
		ctx.Status(http.StatusInternalServerError).Send("err saving file")
		return
	}

	// update users avatar in db
	if err := m.GetDB().Table("accounts").Where("id = ?", userId).Update("avatar", fullPath).Error; err != nil {
		log.Print(err)
		ctx.SendStatus(fiber.StatusUnauthorized)
		ctx.Send("looks like this account doesn't exist")
		return
	}

	ctx.SendStatus(fiber.StatusAccepted)
}
