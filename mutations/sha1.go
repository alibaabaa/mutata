package mutations

import "crypto/sha1"

// Sha1Mutation transforms data to Sha1
func Sha1Mutation() Registration {
	mutation := func(data []byte) []byte {
		hash := sha1.Sum(data)
		return hash[:]
	}
	return Registration{
		Label:    "sha1",
		Mutation: mutation}
}
