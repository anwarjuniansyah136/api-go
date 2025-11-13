package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}

func CheckPassword(hashedPassword, plainPassword string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}