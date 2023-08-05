package strx

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string, cost int) string {
	str, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	if err != nil {
		panic(err)
	}
	return string(str)
}
