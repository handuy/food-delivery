package common

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password []byte) (string, error) {
	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(userPass, submitPass []byte) error {
	return bcrypt.CompareHashAndPassword(userPass, submitPass)
}
