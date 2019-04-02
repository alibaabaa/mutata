package mutations

import "golang.org/x/crypto/bcrypt"

// BCryptMutation transforms data to BCrypt
func BCryptMutation() Registration {
	mutation := func(data []byte) []byte {
		if hash, err := bcrypt.GenerateFromPassword(data, 12); err != nil {
			panic(err)
		} else {
			return hash
		}
	}
	return Registration{
		Label:    "bcrypt",
		Mutation: mutation}
}
