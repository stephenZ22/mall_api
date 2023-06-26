package lib

import "golang.org/x/crypto/bcrypt"

func ComparePassword(hashedPw, Pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(Pw))

	if err != nil {
		return false
	}

	return true
}

func GenerateUserPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return []byte(""), nil
	}

	return hash, nil
}
