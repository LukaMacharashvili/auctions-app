package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 5)

	return string(bytes)
}
