package mutations

import "crypto/sha256"

// Sha256Mutation transforms data to Sha256
func Sha256Mutation() Registration {
	mutation := func(data []byte) []byte {
		hash := sha256.Sum256(data)
		return hash[:]
	}
	return Registration{
		Label:    "sha256",
		Mutation: mutation}
}
