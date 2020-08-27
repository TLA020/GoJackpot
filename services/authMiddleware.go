package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"goprac/models"
	"log"
	"os"
	"strings"
)

var JwtAuthentication = func(c *fiber.Ctx) {

	allowAnonymous := []string{"/api/v1/account/register", "/api/v1/account/login"}

	for _, value := range allowAnonymous {
		if value == c.Path() {
			log.Printf("Public route hit")
			c.Next()
			return
		}
	}

	tokenHeader := c.Get("Authorization")

	if tokenHeader == "" {
		c.Status(401).Send("No token found in header")
		return
	}

	// Format should be like ->  `Bearer {token-body}`, check if  matches this requirement
	tokenParts := strings.Split(tokenHeader, " ")
	if len(tokenParts) != 2 {
		c.Status(401).Send("Token invalid")
		return
	}

	tokenPart := tokenParts[1]
	claimsToken := &models.Token{}

	token, err := jwt.ParseWithClaims(tokenPart, claimsToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		// unable to parse token
		c.Status(401).Send("Parsing token failed")
		return
	}

	if token == nil || !token.Valid {
		c.Status(401).Send("Token expired or invalid")
		return
	}

	// add claims to context.
	c.Locals("user", claimsToken.UserId)
	c.Next()
	return
}
