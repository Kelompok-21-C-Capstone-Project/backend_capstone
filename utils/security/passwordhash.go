package security

import (
	"backend_capstone/services/user"

	"golang.org/x/crypto/bcrypt"
)

type passwordHash struct {
}

func NewPasswordHash() user.PasswordHash {
	var passwordhash *passwordHash
	return passwordhash
}
func (ph *passwordHash) Hash(password string) (hash string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return
	}
	return string(bytes), err
}
func (ph *passwordHash) CheckPassword(password string, hash string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return
	}
	return
}
