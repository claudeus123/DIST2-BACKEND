package utils

import (
	"crypto/rand"
	"math/big"
	"strings"
	// "fmt"
)

const charset = "!@#$%^&*()_+0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
func GeneratePassword(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	var passwordBuilder strings.Builder
	charsetLen := big.NewInt(int64(len(charset)))

	for _, b := range randomBytes {
		idx := new(big.Int).SetBytes([]byte{b})
		idx.Mod(idx, charsetLen)
		passwordBuilder.WriteByte(charset[idx.Int64()])
	}

	return passwordBuilder.String(), nil
}