package bc

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pass string) string {
	salt := 8
	bytePass := []byte(pass)

	hashPass, err := bcrypt.GenerateFromPassword(bytePass, salt)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password.")
	}

	return string(hashPass)
}

func ComparePass(h, p string) bool {
	byteHash, bytePass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePass)

	return err == nil
}
