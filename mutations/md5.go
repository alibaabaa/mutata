package mutations

import "crypto/md5"

// Md5Mutation transforms data to MD5
func Md5Mutation() Registration {
	mutation := func(data []byte) []byte {
		hash := md5.Sum(data)
		return hash[:]
	}
	return Registration{
		Label:    "md5",
		Mutation: mutation}
}
