package hexbyte

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomHexString(numBytes int) string {
	randomBytes, err := GenerateRandomBytes(numBytes)

	if err != nil {
		panic(err)
	}

	return BytesToHex(randomBytes)
}

func GenerateRandomBytes(numBytes int) ([]byte, error) {
	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	return randomBytes, nil
}

func HexToBytes(hexString string) []byte {
	bytes, err := hex.DecodeString(hexString)

	if err != nil {
		panic(err)
	}

	return bytes
}

func BytesToHex(data []byte) string {
	return hex.EncodeToString(data)
}
