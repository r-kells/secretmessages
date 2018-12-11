package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"secretmessages/pkg/crypto"
)

func Encrypt(key crypto.Secret, msg crypto.Message) (secretMsg string, err error) {

	block, err := aes.NewCipher(key.Byte())
	if err != nil {
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	cipherText := make([]byte, aes.BlockSize+len(msg))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], msg.Byte())

	//returns to base64 encoded string
	secretMsg = base64.URLEncoding.EncodeToString(cipherText)
	return
}

func Decrypt(key crypto.Secret, secretMsg crypto.Message) (msg string, err error) {
	cipherText, err := base64.URLEncoding.DecodeString(secretMsg.String())
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key.Byte())
	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		return "", crypto.ErrSecretTooShort
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	msg = string(cipherText)
	return
}
