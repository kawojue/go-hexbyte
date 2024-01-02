package hexbyte

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateRandomHexString generates a random hex string of the specified number of bytes.
// It uses crypto/rand for secure random number generation.
//
// Example:
//
//	randomHex := GenerateRandomHexString(16) // Generates a 32-character random hex string (16 bytes)
func GenerateRandomHexString(numBytes int) string {
	randomBytes, err := GenerateRandomBytes(numBytes)

	if err != nil {
		panic(err)
	}

	return BytesToHex(randomBytes)
}

// GenerateRandomBytes generates a slice of random bytes of the specified length.
// It uses crypto/rand for secure random number generation.
//
// Example:
//
//	randomBytes, err := GenerateRandomBytes(8) // Generates a slice of 8 random bytes
//	if err != nil {
//	    panic(err)
//	}
func GenerateRandomBytes(numBytes int) ([]byte, error) {
	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	return randomBytes, nil
}

// HexToBytes converts a hex string to its byte representation.
//
// Example:
//
//	hexString := "1a2b3c"
//	byteRepresentation := HexToBytes(hexString) // Converts "1a2b3c" to []byte{0x1a, 0x2b, 0x3c}
func HexToBytes(hexString string) []byte {
	bytes, err := hex.DecodeString(hexString)

	if err != nil {
		panic(err)
	}

	return bytes
}

// BytesToHex encodes a byte slice to a hex string.
//
// Example:
//
//	data := []byte{0x1a, 0x2b, 0x3c}
//	hexString := BytesToHex(data) // Converts []byte{0x1a, 0x2b, 0x3c} to "1a2b3c"
func BytesToHex(data []byte) string {
	return hex.EncodeToString(data)
}
