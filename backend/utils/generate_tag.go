package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateTag()(string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	const TAGLENGTH = 4
	b := make([]byte, TAGLENGTH)

	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[n.Int64()]
	}

	return string(b), nil
}