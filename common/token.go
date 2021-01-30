package common

import (
	"errors"
	"log"
	"strings"
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

func GetUserIDFromToken(token string, tokenSecret string) (string, error) {
	var userID string
	splitToken := strings.Split(token, "Bearer ")[1]

	claims := jwt_lib.MapClaims{}

	tkn, err := jwt_lib.ParseWithClaims(splitToken, claims, func(token *jwt_lib.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})

	if err != nil {
		log.Println(err)
		return userID, err
	}

	if !tkn.Valid {
		return userID, errors.New("Token không hợp lệ")
	}

	for k, v := range claims {
		if k == "ID" {
			userID = v.(string)
		}
	}

	if userID == "" {
		return userID, errors.New("Token không hợp lệ")
	}

	return userID, nil
}
