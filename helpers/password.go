package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 8)
	return string(hash), err
}

func ComparePassword(dbPass, userPass string) bool {
	hash, pass := []byte(dbPass), []byte(userPass)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}
