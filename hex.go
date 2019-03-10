package main

import (
	"encoding/hex"
)

// HexTransformer to MD5
type HexTransformer struct{}

// Transform will transform data to MD5
func (m HexTransformer) Transform(data []byte) []byte {
	return []byte(hex.EncodeToString(data))
}
