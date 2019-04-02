package mutations

import "encoding/base64"

// Base64Mutation transforms data to base64
func Base64Mutation() Registration {
	mutation := func(data []byte) []byte {
		return []byte(base64.StdEncoding.EncodeToString(data))
	}
	return Registration{
		Label:    "bcrypt",
		Mutation: mutation}
}
