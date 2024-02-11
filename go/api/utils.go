package api

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hashPass(userPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(userPassword), 14)
	if err != nil {
		return "", fmt.Errorf("error in Hashing password")
	}
	return string(bytes), nil
}

func CheckPassword(password string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, fmt.Errorf("password is incorrect")
	}
	return true, nil
}
