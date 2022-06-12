package util

import (
	"golang.org/x/crypto/bcrypt"
)

func Hasher(textforhash string) (string, error) {
	hashedText, err := bcrypt.GenerateFromPassword([]byte(textforhash), 14)
	return string(hashedText), err
}
