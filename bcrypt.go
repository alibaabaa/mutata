package main

import "golang.org/x/crypto/bcrypt"

// BCryptTransformer to MD5
type BCryptTransformer struct{}

// Transform will transform data to BCrypt
func (b BCryptTransformer) Transform(data []byte) []byte {
	if hash, err := bcrypt.GenerateFromPassword(data, 12); err != nil {
		panic(err)
	} else {
		return hash
	}
}
