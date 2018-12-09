package crypto

import (
	"errors"
)

const LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var ErrSecretTooShort = errors.New("secret is too short!")

// Message to encrypt or decrypt.
type Message string

// Secret to be used to salt certain algorithms.
type Secret string

func (s Secret) String() string {
	return string(s)
}
func (s Secret) Byte() []byte {
	return []byte(s)
}
func (m Message) String() string {
	return string(m)
}
func (m Message) Byte() []byte {
	return []byte(m)
}
