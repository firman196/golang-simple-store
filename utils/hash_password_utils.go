package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func CheckPasswordMatch(hashPassword string, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(currPassword))

	return err == nil
}
