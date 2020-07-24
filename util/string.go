package util

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
	"strconv"
)

// RandomString - generate a random string
func RandomString(length int) string {
	const base = 36
	size := big.NewInt(base)
	n := make([]byte, length)
	for i := range n {
		c, _ := rand.Int(rand.Reader, size)
		n[i] = strconv.FormatInt(c.Int64(), base)[0]
	}
	return string(n)
}

// Comprehension - remove blank strings in slices
func Comprehension(slice []string) []string {
	retSlice := make([]string, 0)
	for _, str := range slice {
		if str != "" {
			retSlice = append(retSlice, str)
		}
	}
	return retSlice
}

// B64Encode - base64 encode
func B64Encode(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}

// B64Decode - base64 decode
func B64Decode(data string) string {
	bytes, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
