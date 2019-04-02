package mutations

import (
	"encoding/hex"
)

// HexMutation transforms data to MD5
func HexMutation() Registration {
	mutation := func(data []byte) []byte {
		return []byte(hex.EncodeToString(data))
	}
	return Registration{
		Label:    "hex",
		Mutation: mutation}
}
