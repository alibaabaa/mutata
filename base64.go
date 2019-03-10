package main

import "encoding/base64"

// Base64Transformer to MD5
type Base64Transformer struct{}

// Transform will encode data to base64
func (m Base64Transformer) Transform(data []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(data))
}
