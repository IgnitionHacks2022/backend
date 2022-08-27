package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var JWTLifetime = 4096
var JWTSecret = os.Getenv("JWT_SECRET")

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

func GenerateToken(id uint) (string, error) {
	expire := time.Now().Add(time.Duration(JWTLifetime) * time.Minute)
	claim := &Claims{UserID: id, StandardClaims: jwt.StandardClaims{ExpiresAt: expire.Unix()}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
