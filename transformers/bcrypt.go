package transformers

import "golang.org/x/crypto/bcrypt"

// BCrypt will transform data to BCrypt
func BCrypt(data []byte) []byte {
	if hash, err := bcrypt.GenerateFromPassword(data, 12); err != nil {
		panic(err)
	} else {
		return hash
	}
}
