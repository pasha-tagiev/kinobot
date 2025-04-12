package random

import (
	"crypto/rand"
	"strings"
)

const DefaultAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_-"

func CryptoText(n int, alphabet string) string {
	buffer := make([]byte, n)
	rand.Read(buffer)

	builder := strings.Builder{}
	builder.Grow(n)

	for _, b := range buffer {
		i := int(b) % len(alphabet)
		builder.WriteByte(alphabet[i])
	}

	return builder.String()
}
