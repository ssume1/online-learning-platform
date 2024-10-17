package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateJWT generate a JWT claims token
func GenerateJWT(email string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.MapClaims{
		"email": email,
		"exp":   expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
