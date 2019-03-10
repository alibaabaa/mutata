package main

import "crypto/md5"

// Md5Transformer to MD5
type Md5Transformer struct{}

// Transform will transform data to MD5
func (m Md5Transformer) Transform(data []byte) []byte {
	hash := md5.Sum(data)
	return hash[:]
}
