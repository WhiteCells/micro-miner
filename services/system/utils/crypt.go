package utils

import "golang.org/x/crypto/bcrypt"

func CryptoPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ValidPassword(userPassword string, inPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(inPassword))
}
