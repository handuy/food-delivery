package common

import (
	"log"
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
)

func NewToken(userID string, tokenSecret string) (string, error) {
	// Create the token
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims = jwt_lib.MapClaims{
		"ID":  userID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}
