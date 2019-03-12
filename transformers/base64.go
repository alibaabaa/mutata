package transformers

import "encoding/base64"

// Base64 to MD5
func Base64(in []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(in))
}
