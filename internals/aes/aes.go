package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// Encrypts the given plain text (bytes). Returns the cipher text if it succeeds in encryting.
// Otherwise error is returned.
func Encrypt(s string) (string, error) {
	key := []byte("2Vzm2hqw9xztBTnjEJRPY46JSUfxg2sE")

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(s), nil)), nil
}

// Decrypts the given cipher text (bytes). Returns the plain text if it succeeds in encryting.
// Otherwise error is returned.
func Decrypt(s string) (string, error) {
	key := []byte("2Vzm2hqw9xztBTnjEJRPY46JSUfxg2sE")
	ciphertext, err := base64.URLEncoding.DecodeString(s)
	//fmt.Println("ciphertext....")

	if err != nil {
		return "", err
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	val, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	//fmt.Println(val)
	return string(val), nil
}
