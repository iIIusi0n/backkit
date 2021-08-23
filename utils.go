package backkit

import (
	"encoding/hex"
	"math/rand"
)

// Generate random string.
func GetRandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// Change byte array to hex string.
func ByteToHexString(b [32]byte) string {
	return hex.EncodeToString(b[:])
}

// Change byte array to hex string.
func HexStringToByte(str string) []byte {
	b, _ := hex.DecodeString(str)
	return b
}
