package utils

import "golang.org/x/crypto/bcrypt"

func HashString(s string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyHash(hash, plainText string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainText))
}
