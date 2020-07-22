package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash For Hashed a String
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckHash For Check Hashed String
func CheckHash(password, hash string) bool {
	crypt := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return crypt == nil
}
