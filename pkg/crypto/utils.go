package crypto

import "math/rand"

// RandStringBytes generates a random string for use as a cipher secret key.
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = LetterBytes[rand.Intn(len(LetterBytes))]
	}
	return string(b)
}
