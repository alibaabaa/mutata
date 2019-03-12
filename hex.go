package main

import (
	"encoding/hex"
)

// Hex will transform data to MD5
func Hex(in []byte) []byte {
	return []byte(hex.EncodeToString(in))
}
