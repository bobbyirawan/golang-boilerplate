package utils

import "golang.org/x/crypto/bcrypt"

func CheckPassword(userPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
}
