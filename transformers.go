package main

import "github.com/alibaabaa/mutata/transformers"

// Transform transforms data from one format to another
type transform func(in []byte) []byte

func getTransformer(transformer string) transform {
	if transformer == "base64" {
		return transformers.Base64
	}
	if transformer == "bcrypt" {
		return transformers.BCrypt
	}
	if transformer == "hex" {
		return transformers.Hex
	}
	if transformer == "md5" {
		return transformers.Md5
	}

	panic("No matching transformer found for " + transformer)
}
