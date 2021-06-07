package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{

			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	return token.SignedString(jwtKey)

}

func ValidateJWT(incJWT string) error {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(incJWT, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !token.Valid {
		return errors.New("invalid token")
	}

	return err

}
