package main

// Transform transforms data from one format to another
type transform func(in []byte) []byte

func getTransformer(transformer string) transform {
	if transformer == "base64" {
		return Base64
	}
	if transformer == "bcrypt" {
		return BCrypt
	}
	if transformer == "hex" {
		return Hex
	}
	if transformer == "md5" {
		return Md5
	}

	panic("No matching transformer found for " + transformer)
}
