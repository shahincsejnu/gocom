package token

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userId string) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}
