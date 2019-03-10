package main

// Transformer transforms data from one format to another
type Transformer interface {
	Transform([]byte) []byte
}

func getTransformer(transformer string) Transformer {
	if transformer == "base64" {
		return Base64Transformer{}
	}
	if transformer == "bcrypt" {
		return BCryptTransformer{}
	}
	if transformer == "hex" {
		return HexTransformer{}
	}
	if transformer == "md5" {
		return Md5Transformer{}
	}

	panic("No matching transformer found for " + transformer)
}
