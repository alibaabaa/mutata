package main

import "crypto/md5"

// Md5 will transform data to MD5
func Md5(data []byte) []byte {
	hash := md5.Sum(data)
	return hash[:]
}
